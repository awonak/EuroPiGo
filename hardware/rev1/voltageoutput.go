package rev1

import (
	"fmt"
	"time"

	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/units"
)

const (
	// Manually calibrated to best match expected voltages. Additional info:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedOffset = 0
	// The default pwmGroup Top of MaxUint16 caused noisy output. Dropping this down to a 8bit value resulted in much smoother cv output.
	CalibratedTop = 0xff - CalibratedOffset

	MaxOutputVoltage = 10.0
	MinOutputVoltage = 0.0
)

// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
// Set a period of 500ns.
var defaultPeriod time.Duration = time.Nanosecond * 500

// voltageoutput is struct for interacting with the CV/VOct voltage output jacks.
type voltageoutput struct {
	pwm PWMProvider
}

var (
	// static check
	_ hal.VoltageOutput = &voltageoutput{}
	// silence linter
	_ = newVoltageOuput
)

type PWMProvider interface {
	Configure(config hal.VoltageOutputConfig) error
	Set(v float32)
	Get() float32
	MinVoltage() float32
	MaxVoltage() float32
}

// NewOutput returns a new Output interface.
func newVoltageOuput(pwm PWMProvider) hal.VoltageOutput {
	o := &voltageoutput{
		pwm: pwm,
	}
	err := o.Configure(hal.VoltageOutputConfig{
		Period: defaultPeriod,
		Calibration: envelope.NewMap32([]envelope.MapEntry[float32, uint16]{
			{
				Input:  MinOutputVoltage,
				Output: CalibratedTop,
			},
			{
				Input:  MaxOutputVoltage,
				Output: CalibratedOffset,
			},
		}),
	})
	if err != nil {
		panic(fmt.Errorf("configuration error: %v", err.Error()))
	}

	return o
}

// Configure updates the device with various configuration parameters
func (o *voltageoutput) Configure(config hal.VoltageOutputConfig) error {
	if err := o.pwm.Configure(config); err != nil {
		return err
	}

	return nil
}

// SetVoltage sets the current output voltage within a range of 0.0 to 10.0.
func (o *voltageoutput) SetVoltage(v float32) {
	o.pwm.Set(v)
}

// SetCV sets the current output voltage based on a CV value
func (o *voltageoutput) SetCV(cv units.CV) {
	o.SetVoltage(cv.ToVolts())
}

// SetBipolarCV sets the current output voltage based on a BipolarCV value
func (o *voltageoutput) SetBipolarCV(cv units.BipolarCV) {
	o.SetVoltage(cv.ToVolts())
}

// SetCV sets the current output voltage based on a V/Octave value
func (o *voltageoutput) SetVOct(voct units.VOct) {
	o.SetVoltage(voct.ToVolts())
}

// Voltage returns the current voltage
func (o *voltageoutput) Voltage() float32 {
	return o.pwm.Get()
}

// MinVoltage returns the minimum voltage this device will output
func (o *voltageoutput) MinVoltage() float32 {
	return o.pwm.MinVoltage()
}

// MaxVoltage returns the maximum voltage this device will output
func (o *voltageoutput) MaxVoltage() float32 {
	return o.pwm.MaxVoltage()
}
