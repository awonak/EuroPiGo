package common

import (
	"errors"

	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/units"
)

// Analog is a struct for handling the reading of analogue control voltage.
// The analogue input allows you to 'read' CV from anywhere between 0 and 12V.
type Analoginput struct {
	adc     ADCProvider
	samples int
	cal     envelope.Map[uint16, float32]
}

var (
	// static check
	_ hal.AnalogInput = (*Analoginput)(nil)
	// silence linter
	_ = NewAnalogInput
)

type ADCProvider interface {
	Get(samples int) uint16
}

// NewAnalogInput creates a new Analog Input
func NewAnalogInput(adc ADCProvider, initialConfig hal.AnalogInputConfig) *Analoginput {
	if adc == nil {
		return nil
	}
	return &Analoginput{
		adc:     adc,
		samples: initialConfig.Samples,
		cal:     initialConfig.Calibration,
	}
}

// Configure updates the device with various configuration parameters
func (a *Analoginput) Configure(config hal.AnalogInputConfig) error {
	if config.Samples == 0 {
		return errors.New("samples must be non-zero")
	}

	if config.Calibration != nil {
		a.cal = config.Calibration
	}

	a.samples = config.Samples
	return nil
}

// ReadRawVoltage returns the current smoothed value from the analog input device.
func (a *Analoginput) ReadRawVoltage() uint16 {
	return a.adc.Get(a.samples)
}

// ReadVoltage returns the current percentage read between 0.0 and 1.0.
func (a *Analoginput) Percent() float32 {
	return a.ReadVoltage() / a.cal.OutputMaximum()
}

// ReadVoltage returns the current read voltage between 0.0 and 10.0 volts.
func (a *Analoginput) ReadVoltage() float32 {
	rawVoltage := a.ReadRawVoltage()
	return a.cal.Remap(rawVoltage)
}

// ReadCV returns the current read voltage as a CV value.
func (a *Analoginput) ReadCV() units.CV {
	v := a.ReadVoltage()
	// CV is ranged over 0.0V .. +5.0V and stores the values as a normalized
	// version (0.0 .. +1.0), so to convert our input voltage to that, we just
	// normalize the voltage (divide it by 5) and clamp the result.
	return clamp.Clamp(units.CV(v/5.0), 0.0, 1.0)
}

func (a *Analoginput) ReadBipolarCV() units.BipolarCV {
	v := a.ReadVoltage()
	// BipolarCV is ranged over -5.0V .. +5.0V and stores the values as a normalized
	// version (-1.0 .. +1.0), so to convert our input voltage to that, we just
	// normalize the voltage (divide it by 5) and clamp the result.
	return clamp.Clamp(units.BipolarCV(v/5.0), -1.0, 1.0)
}

// ReadCV returns the current read voltage as a V/Octave value.
func (a *Analoginput) ReadVOct() units.VOct {
	return units.VOct(a.ReadVoltage())
}

// MinVoltage returns the minimum voltage that that input can ever read by this device
func (a *Analoginput) MinVoltage() float32 {
	return a.cal.OutputMinimum()
}

// MaxVoltage returns the maximum voltage that the input can ever read by this device
func (a *Analoginput) MaxVoltage() float32 {
	return a.cal.OutputMaximum()
}
