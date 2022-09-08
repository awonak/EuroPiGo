package europi

import (
	"fmt"
	"machine"
)

const (
	// Manually calibrated to best match expected voltages. Additional info:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedOffset = 32
	// The default PWM Top of MaxUint16 caused noisy output. Dropping this down to a 12bit value resulted in much smoother cv output.
	CalibratedTop = 0xfff - CalibratedOffset
)

// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
// For a frequency of 1mHz, we must set a period of 1000ns.
var defaultPeriod uint64 = 1000

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

type Outputer interface {
	Get() (value uint32)
	Voltage(v float32)
	On()
	Off()
}

type Output struct {
	pwm PWMer
	pin machine.Pin
	ch  uint8
}

func NewOutput(pin machine.Pin, pwm PWMer) *Output {
	err := pwm.Configure(machine.PWMConfig{
		Period: defaultPeriod,
	})
	if err != nil {
		fmt.Println("pwm Configure error: ", err.Error())
	}

	pwm.SetTop(CalibratedTop)

	ch, err := pwm.Channel(pin)
	if err != nil {
		fmt.Println("pwm Channel error: ", err.Error())
	}

	return &Output{pwm, pin, ch}
}

func (o *Output) Get() uint32 {
	return o.pwm.Get(o.ch)
}

func (o *Output) Voltage(v float32) {
	// TODO: boundary check
	invertedCv := (v / MaxVoltage) * float32(o.pwm.Top())
	cv := (float32(o.pwm.Top()) - invertedCv) - CalibratedOffset
	o.pwm.Set(o.ch, uint32(cv))
}

func (o *Output) On() {
	o.pwm.Set(o.ch, o.pwm.Top())
}

func (o *Output) Off() {
	o.pwm.Set(o.ch, 0)
}
