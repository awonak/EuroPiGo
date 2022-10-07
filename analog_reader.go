package europi

import (
	"machine"
	"math"
)

const (
	// Calibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedMinAI = 300
	CalibratedMaxAI = 44009

	DefaultSamples = 1000
)

var (
	AI AnalogReader
	K1 AnalogReader
	K2 AnalogReader
)

func init() {
	machine.InitADC()
	AI = newAI(machine.ADC0)
	K1 = newKnob(machine.ADC1)
	K2 = newKnob(machine.ADC2)
}

// AnalogReader is an interface for common analog read methods for knobs and cv input.
type AnalogReader interface {
	Samples(samples uint16)
	ReadVoltage() float32
	Percent() float32
	Range(steps uint16) uint16
}

// A struct for handling the reading of analogue control voltage.
// The analogue input allows you to 'read' CV from anywhere between 0 and 12V.
type analogInput struct {
	machine.ADC
	samples uint16
}

// newAI creates a new AnalogInput.
func newAI(pin machine.Pin) *analogInput {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &analogInput{ADC: adc, samples: DefaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (a *analogInput) Samples(samples uint16) {
	a.samples = samples
}

// Percent return the percentage of the input's current relative range as a float between 0.0 and 1.0.
func (a *analogInput) Percent() float32 {
	return float32(a.read()) / CalibratedMaxAI
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (a *analogInput) ReadVoltage() float32 {
	return a.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the analog input.
func (a *analogInput) Range(steps uint16) uint16 {
	return uint16(a.Percent() * float32(steps))
}

func (a *analogInput) read() uint16 {
	var sum int
	for i := 0; i < int(a.samples); i++ {
		sum += Clamp(int(a.Get())-CalibratedMinAI, 0, CalibratedMaxAI)
	}
	return uint16(sum / int(a.samples))
}

// A struct for handling the reading of knob voltage and position.
type knob struct {
	machine.ADC
	samples uint16
}

// newKnob creates a new Knob struct.
func newKnob(pin machine.Pin) *knob {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &knob{ADC: adc, samples: DefaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (k *knob) Samples(samples uint16) {
	k.samples = samples
}

// Percent return the percentage of the knob's current relative range as a float between 0.0 and 1.0.
func (k *knob) Percent() float32 {
	return 1 - float32(k.read())/math.MaxUint16
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (k *knob) ReadVoltage() float32 {
	return k.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the knob's position.
func (k *knob) Range(steps uint16) uint16 {
	return uint16(k.Percent() * float32(steps))
}

func (k *knob) read() uint16 {
	var sum int
	for i := 0; i < int(k.samples); i++ {
		sum += int(k.Get())
	}
	return uint16(sum / int(k.samples))
}
