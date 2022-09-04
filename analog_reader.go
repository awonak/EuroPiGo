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

	DefaultSamples = 100
)

func init() {
	machine.InitADC()
}

type AnalogReader interface {
	SetSamples(samples uint16)
	ReadVoltage() uint16
	Percent() float32
	Range(steps uint16) uint16
}

type AnalogInput struct {
	machine.ADC
	samples uint16
}

func NewAI(pin machine.Pin) *AnalogInput {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &AnalogInput{ADC: adc, samples: DefaultSamples}
}

func (a *AnalogInput) Samples(samples uint16) {
	a.samples = samples
}

func (a *AnalogInput) Percent() float32 {
	return float32(a.read()) / CalibratedMaxAI
}

func (a *AnalogInput) ReadVoltage() uint16 {
	return uint16(a.Percent() * MaxVoltage)
}

func (a *AnalogInput) Range(steps uint16) uint16 {
	return uint16(a.Percent() * float32(steps))
}

func (a *AnalogInput) read() uint16 {
	var i uint16
	sum := 0
	for i = 0; i < a.samples; i++ {
		sum += int(Clamp(int(a.Get())-CalibratedMinAI, 0, CalibratedMaxAI))
	}
	return uint16(sum / int(a.samples))
}

type Knob struct {
	machine.ADC
	samples uint16
}

func NewKnob(pin machine.Pin) *Knob {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &Knob{ADC: adc, samples: DefaultSamples}
}

func (k *Knob) Samples(samples uint16) {
	k.samples = samples
}

func (k *Knob) Percent() float32 {
	return 1 - float32(k.read())/math.MaxUint16
}

func (k *Knob) ReadVoltage() uint16 {
	return uint16(k.Percent() * MaxVoltage)
}

func (k *Knob) Range(steps uint16) uint16 {
	return uint16(k.Percent() * float32(steps))
}

func (k *Knob) read() uint16 {
	var i uint16
	sum := 0
	for i = 0; i < k.samples; i++ {
		sum += int(k.Get())
	}
	return uint16(sum / int(k.samples))
}
