package rev1

import (
	"fmt"
	"time"

	"github.com/heucuva/europi/clamp"
	"github.com/heucuva/europi/internal/hardware/hal"
	"github.com/heucuva/europi/units"
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
	pwm pwmProvider
	ofs uint16
}

// NewOutput returns a new Output interface.
func newVoltageOuput(pwm pwmProvider) hal.VoltageOutput {
	o := &voltageoutput{
		pwm: pwm,
	}
	err := o.Configure(hal.VoltageOutputConfig{
		Period: defaultPeriod,
		Offset: CalibratedOffset,
		Top:    CalibratedTop,
	})
	if err != nil {
		panic(fmt.Errorf("configuration error: %v", err.Error()))
	}

	return o
}

type pwmProvider interface {
	Configure(config hal.VoltageOutputConfig) error
	Set(v float32, ofs uint16)
	Get() float32
}

func (o *voltageoutput) Configure(config hal.VoltageOutputConfig) error {
	if err := o.pwm.Configure(config); err != nil {
		return err
	}

	o.ofs = config.Offset

	return nil
}

// SetVoltage sets the current output voltage within a range of 0.0 to 10.0.
func (o *voltageoutput) SetVoltage(v float32) {
	v = clamp.Clamp(v, MinOutputVoltage, MaxOutputVoltage)
	o.pwm.Set(v/MaxOutputVoltage, o.ofs)
}

// SetCV sets the current output voltage based on a CV value
func (o *voltageoutput) SetCV(cv units.CV) {
	o.SetVoltage(cv.ToVolts())
}

// SetCV sets the current output voltage based on a V/Octave value
func (o *voltageoutput) SetVOct(voct units.VOct) {
	o.SetVoltage(voct.ToVolts())
}

// Voltage returns the current voltage
func (o *voltageoutput) Voltage() float32 {
	return o.pwm.Get() * MaxOutputVoltage
}

func (o *voltageoutput) MinVoltage() float32 {
	return MinOutputVoltage
}

func (o *voltageoutput) MaxVoltage() float32 {
	return MaxOutputVoltage
}
