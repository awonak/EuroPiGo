package europi

import (
	"machine"
)

const (
	// Manually calibrated to best match expected voltages. Additional info:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	calibratedOffset = 0
	// The default PWM Top of MaxUint16 caused noisy output. Dropping this down to a 8bit value resulted in much smoother cv output.
	calibratedTop = 0xff - calibratedOffset
)

var (
	// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
	// Set a period of 500ns.
	defaultPeriod uint64 = 500
)

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

// output is struct for interacting with the cv output jacks.
type output struct {
	pwm PWMer
	pin machine.Pin
	ch  uint8
}

func newOutput(pin machine.Pin, pwm PWMer) *output {
	err := pwm.Configure(machine.PWMConfig{
		Period: defaultPeriod,
	})
	if err != nil {
		panic("PWM Configure error")
	}

	pwm.SetTop(calibratedTop)

	ch, err := pwm.Channel(pin)
	if err != nil {
		panic("PWM Channel error")
	}

	return &output{pwm, pin, ch}
}

// Get returns the current set voltage in the range of 0 to pwm.Top().
func (o *output) Get() uint32 {
	return o.pwm.Get(o.ch)
}

// Voltage sets the current output voltage within a range of 0.0 to 10.0.
func (o *output) Voltage(v float32) {
	v = Clamp(v, MinVoltage, MaxVoltage)
	invertedCv := (v / MaxVoltage) * float32(o.pwm.Top())
	// cv := (float32(o.pwm.Top()) - invertedCv) - CalibratedOffset
	cv := float32(invertedCv) - calibratedOffset
	o.pwm.Set(o.ch, uint32(cv))
}

// On sets the current voltage high at 10.0v.
func (o *output) On() {
	o.pwm.Set(o.ch, o.pwm.Top())
}

// Off sets the current voltage low at 0.0v.
func (o *output) Off() {
	o.pwm.Set(o.ch, 0)
}
