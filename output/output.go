package output

import (
	"log"
	"machine"
	"math"
	"runtime/interrupt"
	"runtime/volatile"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

const (
	// Manually calibrated to best match expected voltages. Additional info:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedOffset = 0
	// The default PWM Top of MaxUint16 caused noisy output. Dropping this down to a 8bit value resulted in much smoother cv output.
	CalibratedTop = 0xff - CalibratedOffset

	MaxVoltage = 10.0
	MinVoltage = 0.0
)

// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
// Set a period of 500ns.
var defaultPeriod uint64 = 500

// Output is an interface for interacting with the cv output jacks.
type Output interface {
	Get() uint32
	SetVoltage(v float32)
	SetCV(cv units.CV)
	SetVOct(voct units.VOct)
	Set(v bool)
	On()
	Off()
	Voltage() float32
}

// Output is struct for interacting with the cv output jacks.
type output struct {
	pwm PWM
	pin machine.Pin
	ch  uint8
	v   uint32
}

// NewOutput returns a new Output interface.
func NewOutput(pin machine.Pin, pwm PWM) Output {
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

	return &output{pwm, pin, ch, MinVoltage}
}

// Get returns the current set voltage in the range of 0 to pwm.Top().
func (o *output) Get() uint32 {
	state := interrupt.Disable()
	v := o.pwm.Get(o.ch)
	interrupt.Restore(state)
	return v
}

// Set updates the current voltage high (true) or low (false)
func (o *output) Set(v bool) {
	if v {
		o.On()
	} else {
		o.Off()
	}
}

// SetVoltage sets the current output voltage within a range of 0.0 to 10.0.
func (o *output) SetVoltage(v float32) {
	v = europim.Clamp(v, MinVoltage, MaxVoltage)
	invertedCv := (v / MaxVoltage) * float32(o.pwm.Top())
	// cv := (float32(o.pwm.Top()) - invertedCv) - CalibratedOffset
	cv := float32(invertedCv) - CalibratedOffset
	state := interrupt.Disable()
	o.pwm.Set(o.ch, uint32(cv))
	interrupt.Restore(state)
	volatile.StoreUint32(&o.v, math.Float32bits(v))
}

// SetCV sets the current output voltage based on a CV value
func (o *output) SetCV(cv units.CV) {
	o.SetVoltage(cv.ToVolts())
}

// SetCV sets the current output voltage based on a V/Octave value
func (o *output) SetVOct(voct units.VOct) {
	o.SetVoltage(voct.ToVolts())
}

// On sets the current voltage high at 10.0v.
func (o *output) On() {
	volatile.StoreUint32(&o.v, math.Float32bits(MaxVoltage))
	state := interrupt.Disable()
	o.pwm.Set(o.ch, o.pwm.Top())
	interrupt.Restore(state)
}

// Off sets the current voltage low at 0.0v.
func (o *output) Off() {
	volatile.StoreUint32(&o.v, math.Float32bits(MinVoltage))
	state := interrupt.Disable()
	o.pwm.Set(o.ch, 0)
	interrupt.Restore(state)
}

// Voltage returns the current voltage
func (o *output) Voltage() float32 {
	return math.Float32frombits(volatile.LoadUint32(&o.v))
}
