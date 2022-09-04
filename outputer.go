package europi

import (
	"fmt"
	"machine"
	"math"
)

const (
	// CalibratedMaxDuty was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedMaxDuty = 63475
	CalibratedOffset  = math.MaxUint16 - CalibratedMaxDuty
)

var defaultPeriod uint64 = 1e9 / 2000

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

	pwm.Configure(machine.PWMConfig{
		Period: defaultPeriod,
	})

	pwm.SetTop(CalibratedMaxDuty)

	ch, err := pwm.Channel(pin)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &Output{pwm, pin, ch}
}

func (o *Output) Voltage(v float32) {
	// TODO: boundary check
	invertedCv := (v / MaxVoltage) * float32(o.pwm.Top())
	cv := (math.MaxUint16 - invertedCv) - CalibratedOffset
	o.pwm.Set(o.ch, uint32(cv))
}

func (o *Output) On() {
	o.pwm.Set(o.ch, o.pwm.Top())
}

func (o *Output) Off() {
	o.pwm.Set(o.ch, 0)
}
