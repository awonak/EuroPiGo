package europi

import (
	"log"
	"machine"
)

const (
	// Manually calibrated to best match expected voltages. Additional info:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedOffset = 3
	// The default PWM Top of MaxUint16 caused noisy output. Dropping this down to a 8bit value resulted in much smoother cv output.
	CalibratedTop = 0xff - CalibratedOffset
)

// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
// Set a period of 500ns.
var defaultPeriod uint64 = 500

// PWMer is an interface for interacting with a machine.pwmGroup
type PWMer interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	SetTop(top uint32)
	Get(channel uint8) (value uint32)
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

// Outputer is an interface for interacting with the cv output jacks.
type Outputer interface {
	Get() (value uint32)
	Voltage(v float32)
	On()
	Off()
}

// Outputer is struct for interacting with the cv output jacks.
type Output struct {
	pwm PWMer
	pin machine.Pin
	ch  uint8
}

// NewOutput returns a new Output struct.
func NewOutput(pin machine.Pin, pwm PWMer) *Output {
	err := pwm.Configure(machine.PWMConfig{
		Period: defaultPeriod,
	})
	if err != nil {
		log.Fatal("pwm Configure error: ", err.Error())
	}

	pwm.SetTop(CalibratedTop)

	ch, err := pwm.Channel(pin)
	if err != nil {
		log.Fatal("pwm Channel error: ", err.Error())
	}

	return &Output{pwm, pin, ch}
}

// Get returns the current set voltage in the range of 0 to pwm.Top().
func (o *Output) Get() uint32 {
	return o.pwm.Get(o.ch)
}

// Voltage sets the current output voltage within a range of 0.0 to 10.0.
func (o *Output) Voltage(v float32) {
	v = Clamp(v, MinVoltage, MaxVoltage)
	invertedCv := (v / MaxVoltage) * float32(o.pwm.Top())
	cv := (float32(o.pwm.Top()) - invertedCv) - CalibratedOffset
	o.pwm.Set(o.ch, uint32(cv))
}

// On sets the current voltage high at 10.0v.
func (o *Output) On() {
	o.pwm.Set(o.ch, o.pwm.Top())
}

// Off sets the current voltage low at 0.0v.
func (o *Output) Off() {
	o.pwm.Set(o.ch, 0)
}
