package input

import (
	"machine"
	"math"
)

// A struct for handling the reading of knob voltage and position.
type Knob struct {
	machine.ADC
	samples uint16
}

// NewKnob creates a new Knob struct.
func NewKnob(pin machine.Pin) *Knob {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &Knob{ADC: adc, samples: DefaultSamples}
}

// Samples sets the number of reads for an more accurate average read.
func (k *Knob) Samples(samples uint16) {
	k.samples = samples
}

// Percent return the percentage of the knob's current relative range as a float between 0.0 and 1.0.
func (k *Knob) Percent() float32 {
	return 1 - float32(k.read())/math.MaxUint16
}

// ReadVoltage return the current read voltage between 0.0 and 10.0 volts.
func (k *Knob) ReadVoltage() float32 {
	return k.Percent() * MaxVoltage
}

// Range return a value between 0 and the given steps (not inclusive) based on the range of the knob's position.
func (k *Knob) Range(steps uint16) uint16 {
	return uint16(k.Percent() * float32(steps))
}

func (k *Knob) read() uint16 {
	var sum int
	for i := 0; i < int(k.samples); i++ {
		sum += int(k.Get())
	}
	return uint16(sum / int(k.samples))
}
