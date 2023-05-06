package common

import (
	"fmt"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/units"
)

// VoltageOutput is struct for interacting with the CV/VOct voltage output jacks.
type VoltageOutput struct {
	pwm PWMProvider
}

var (
	// static check
	_ hal.VoltageOutput = (*VoltageOutput)(nil)
	// silence linter
	_ = NewVoltageOuput
)

type PWMProvider interface {
	Configure(config hal.VoltageOutputConfig) error
	Set(v float32)
	Get() float32
	MinVoltage() float32
	MaxVoltage() float32
}

// NewOutput returns a new Output interface.
func NewVoltageOuput(pwm PWMProvider, initialConfig hal.VoltageOutputConfig) *VoltageOutput {
	o := &VoltageOutput{
		pwm: pwm,
	}
	err := o.Configure(initialConfig)
	if err != nil {
		panic(fmt.Errorf("configuration error: %v", err.Error()))
	}

	return o
}

// Configure updates the device with various configuration parameters
func (o *VoltageOutput) Configure(config hal.VoltageOutputConfig) error {
	if err := o.pwm.Configure(config); err != nil {
		return err
	}

	return nil
}

// SetVoltage sets the current output voltage within a range of 0.0 to 10.0.
func (o *VoltageOutput) SetVoltage(v float32) {
	o.pwm.Set(v)
}

// SetCV sets the current output voltage based on a CV value
func (o *VoltageOutput) SetCV(cv units.CV) {
	o.SetVoltage(cv.ToVolts())
}

// SetBipolarCV sets the current output voltage based on a BipolarCV value
func (o *VoltageOutput) SetBipolarCV(cv units.BipolarCV) {
	o.SetVoltage(cv.ToVolts())
}

// SetCV sets the current output voltage based on a V/Octave value
func (o *VoltageOutput) SetVOct(voct units.VOct) {
	o.SetVoltage(voct.ToVolts())
}

// Voltage returns the current voltage
func (o *VoltageOutput) Voltage() float32 {
	return o.pwm.Get()
}

// MinVoltage returns the minimum voltage this device will output
func (o *VoltageOutput) MinVoltage() float32 {
	return o.pwm.MinVoltage()
}

// MaxVoltage returns the maximum voltage this device will output
func (o *VoltageOutput) MaxVoltage() float32 {
	return o.pwm.MaxVoltage()
}
