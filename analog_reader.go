package europi

import (
	"machine"
	"math"
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
}

type analogInput struct {
	adc     machine.ADC
	samples uint16
}

func newAnalogInput(pin machine.Pin) *analogInput {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &analogInput{adc: adc, samples: defaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (a *analogInput) Samples(samples uint16) {
	a.samples = samples
}

// Percent return the percentage of the input's current relative range as a float between 0.0 and 1.0.
func (a *analogInput) Percent() float32 {
	return float32(a.read()) / calibratedMaxAI
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
		sum += Clamp(int(a.adc.Get())-calibratedMinAI, 0, calibratedMaxAI)
	}
	return uint16(sum / int(a.samples))
}

type knob struct {
	adc     machine.ADC
	samples uint16
}

func newKnob(pin machine.Pin) *knob {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &knob{adc: adc, samples: defaultSamples}
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
		sum += int(k.adc.Get())
	}
	return uint16(sum / int(k.samples))
}
