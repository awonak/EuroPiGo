package europi

import (
	"machine"
)

const (
	// Default number of analog reads to average over.
	defaultSamples = 32

	// Calibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	calibratedMinAI = 300
	calibratedMaxAI = 44009
)

func init() {
	machine.InitADC()
}

// AnalogReader is an interface for common analog read methods for knobs and cv input.
type AnalogReader interface {
	Samples(samples uint16)
	ReadVoltage() float32
	Percent() float32
	Range(steps uint16) uint16
	read() uint16
}

type analogReader struct {
	machine.ADC

	samples uint16
}

func newAnalogReader(pin machine.Pin) *analogReader {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &analogReader{
		ADC:     adc,
		samples: defaultSamples,
	}
}

// Samples sets the number of reads for an more accurate average read.
func (a *analogReader) Samples(samples uint16) {
	a.samples = samples
}

// Percent return the percentage of the input's current relative range as a float between 0.0 and 1.0.
func (a *analogReader) Percent() float32 {
	return float32(a.read()) / calibratedMaxAI
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (a *analogReader) ReadVoltage() float32 {
	return a.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the analog input.
func (a *analogReader) Range(steps uint16) uint16 {
	return uint16(a.Percent() * float32(steps))
}

func (a *analogReader) read() uint16 {
	var sum int
	for i := 0; i < int(a.samples); i++ {
		sum += Clamp(int(a.Get())-calibratedMinAI, 0, calibratedMaxAI)
	}
	return uint16(sum / int(a.samples))
}

type analogInput struct {
	AnalogReader
}

func newAnalogInput(pin machine.Pin) *analogInput {
	return &analogInput{
		newAnalogReader(pin),
	}
}

type knob struct {
	AnalogReader
}

func newKnob(pin machine.Pin) *knob {
	return &knob{
		newAnalogReader(pin),
	}
}
