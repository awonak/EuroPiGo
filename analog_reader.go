package europi

import (
	"machine"
	"math"
)

const (
	// Calibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedMinAI = 384
	CalibratedMaxAI = 44634

	DefaultSamples = 1000
)

func init() {
	machine.InitADC()
}

// AnalogReader is an interface for common analog read methods for knobs and cv input.
type AnalogReader interface {
	Samples(samples int)
	ReadVoltage() float32
	Percent() float32
	Range(steps int) int
}

// A struct for handling the reading of analogue control voltage.
// The analogue input allows you to 'read' CV from anywhere between 0 and 12V.
type AnalogInput struct {
	machine.ADC
	samples int
}

// NewAI creates a new AnalogInput.
func NewAI(pin machine.Pin) *AnalogInput {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &AnalogInput{ADC: adc, samples: DefaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (a *AnalogInput) Samples(samples int) {
	a.samples = samples
}

// Percent return the percentage of the input's current relative range.
func (a *AnalogInput) Percent() float32 {
	return float32(a.read()) / CalibratedMaxAI
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (a *AnalogInput) ReadVoltage() float32 {
	return a.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the analog input.
func (a *AnalogInput) Range(steps int) int {
	return int(a.Percent() * float32(steps))
}

func (a *AnalogInput) read() int {
	sum := 0
	for i := 0; i < a.samples; i++ {
		sum += Clamp(int(a.Get())-CalibratedMinAI, 0, CalibratedMaxAI)
	}
	return sum / a.samples
}

// A struct for handling the reading of knob voltage and position.
type Knob struct {
	machine.ADC
	samples int
}

// NewKnob creates a new Knob struct.
func NewKnob(pin machine.Pin) *Knob {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &Knob{ADC: adc, samples: DefaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (k *Knob) Samples(samples int) {
	k.samples = samples
}

// Percent return the percentage of the knob's current relative range.
func (k *Knob) Percent() float32 {
	return 1 - float32(k.read())/math.MaxUint16
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (k *Knob) ReadVoltage() float32 {
	return k.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the knob's position.
func (k *Knob) Range(steps int) int {
	return int(k.Percent() * float32(steps))
}

func (k *Knob) read() int {
	sum := 0
	for i := 0; i < k.samples; i++ {
		sum += int(k.Get())
	}
	return int(sum / k.samples)
}
