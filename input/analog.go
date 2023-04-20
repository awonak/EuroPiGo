package input

import (
	"machine"
	"runtime/interrupt"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
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

// ReadCV returns the current read voltage as a CV value.
func (a *Analog) ReadCV() units.CV {
	// we can't use a.Percent() here, because we might get over 5.0 volts input
	// just clamp it
	v := a.ReadVoltage()
	return units.CV(europim.Clamp(v/5.0, 0.0, 1.0))
}

// ReadCV returns the current read voltage as a V/Octave value.
func (a *Analog) ReadVOct() units.VOct {
	return units.VOct(a.ReadVoltage())
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the analog input.
func (a *Analog) Range(steps uint16) uint16 {
	return uint16(a.Percent() * float32(steps))
}

func (a *Analog) Choice(numItems int) int {
	return europim.Lerp(a.Percent(), 0, numItems-1)
}

func (a *Analog) read() uint16 {
	var sum int
	state := interrupt.Disable()
	for i := 0; i < int(a.samples); i++ {
		sum += europim.Clamp(int(a.Get())-CalibratedMinAI, 0, CalibratedMaxAI)
	}
	interrupt.Restore(state)
	return uint16(sum / int(a.samples))
}
