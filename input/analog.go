package input

import (
	"machine"

	europiMath "github.com/heucuva/europi/internal/math"
)

const (
	// Calibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedMinAI = 300
	CalibratedMaxAI = 44009

	DefaultSamples = 1000

	MaxVoltage = 10.0
	MinVoltage = 0.0
)

// Analog is a struct for handling the reading of analogue control voltage.
// The analogue input allows you to 'read' CV from anywhere between 0 and 12V.
type Analog struct {
	machine.ADC
	samples uint16
}

// NewAnalog creates a new Analog.
func NewAnalog(pin machine.Pin) *Analog {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &Analog{ADC: adc, samples: DefaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (a *Analog) Samples(samples uint16) {
	a.samples = samples
}

// Percent return the percentage of the input's current relative range as a float between 0.0 and 1.0.
func (a *Analog) Percent() float32 {
	return float32(a.read()) / CalibratedMaxAI
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (a *Analog) ReadVoltage() float32 {
	return a.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the analog input.
func (a *Analog) Range(steps uint16) uint16 {
	return uint16(a.Percent() * float32(steps))
}

func (a *Analog) read() uint16 {
	var sum int
	for i := 0; i < int(a.samples); i++ {
		sum += europiMath.Clamp(int(a.Get())-CalibratedMinAI, 0, CalibratedMaxAI)
	}
	return uint16(sum / int(a.samples))
}
