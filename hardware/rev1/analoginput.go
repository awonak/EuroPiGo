package rev1

import (
	"errors"

	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/units"
)

const (
	// DefaultCalibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	DefaultCalibratedMinAI = 300
	DefaultCalibratedMaxAI = 44009

	DefaultSamples = 1000

	MaxInputVoltage = 10.0
	MinInputVoltage = 0.0
)

// Analog is a struct for handling the reading of analogue control voltage.
// The analogue input allows you to 'read' CV from anywhere between 0 and 12V.
type analoginput struct {
	adc     ADCProvider
	samples int
	cal     envelope.Map[uint16, float32]
}

var (
	// static check
	_ hal.AnalogInput = &analoginput{}
	// silence linter
	_ = newAnalogInput
)

type ADCProvider interface {
	Get(samples int) uint16
}

// newAnalogInput creates a new Analog Input
func newAnalogInput(adc ADCProvider) *analoginput {
	return &analoginput{
		adc:     adc,
		samples: DefaultSamples,
		cal: envelope.NewMap32([]envelope.MapEntry[uint16, float32]{
			{
				Input:  DefaultCalibratedMinAI,
				Output: MinInputVoltage,
			},
			{
				Input:  DefaultCalibratedMaxAI,
				Output: MaxInputVoltage,
			},
		}),
	}
}

// Configure updates the device with various configuration parameters
func (a *analoginput) Configure(config hal.AnalogInputConfig) error {
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
func (a *analoginput) ReadRawVoltage() uint16 {
	return a.adc.Get(a.samples)
}

// ReadVoltage returns the current percentage read between 0.0 and 1.0.
func (a *analoginput) Percent() float32 {
	return a.ReadVoltage() / MaxInputVoltage
}

// ReadVoltage returns the current read voltage between 0.0 and 10.0 volts.
func (a *analoginput) ReadVoltage() float32 {
	rawVoltage := a.ReadRawVoltage()
	return a.cal.Remap(rawVoltage)
}

// ReadCV returns the current read voltage as a CV value.
func (a *analoginput) ReadCV() units.CV {
	v := a.ReadVoltage()
	// CV is ranged over 0.0V .. +5.0V and stores the values as a normalized
	// version (0.0 .. +1.0), so to convert our input voltage to that, we just
	// normalize the voltage (divide it by 5) and clamp the result.
	return clamp.Clamp(units.CV(v/5.0), 0.0, 1.0)
}

func (a *analoginput) ReadBipolarCV() units.BipolarCV {
	v := a.ReadVoltage()
	// BipolarCV is ranged over -5.0V .. +5.0V and stores the values as a normalized
	// version (-1.0 .. +1.0), so to convert our input voltage to that, we just
	// normalize the voltage (divide it by 5) and clamp the result.
	return clamp.Clamp(units.BipolarCV(v/5.0), -1.0, 1.0)
}

// ReadCV returns the current read voltage as a V/Octave value.
func (a *analoginput) ReadVOct() units.VOct {
	return units.VOct(a.ReadVoltage())
}

// MinVoltage returns the minimum voltage that that input can ever read by this device
func (a *analoginput) MinVoltage() float32 {
	return a.cal.OutputMinimum()
}

// MaxVoltage returns the maximum voltage that the input can ever read by this device
func (a *analoginput) MaxVoltage() float32 {
	return a.cal.OutputMaximum()
}
