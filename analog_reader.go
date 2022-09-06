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

type AnalogReader interface {
	Samples(samples int)
	ReadVoltage() float32
	Percent() float32
	Range(steps int) int
}

type AnalogInput struct {
	machine.ADC
	samples int
}

func NewAI(pin machine.Pin) *AnalogInput {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &AnalogInput{ADC: adc, samples: DefaultSamples}
}

func (a *AnalogInput) Samples(samples int) {
	a.samples = samples
}

func (a *AnalogInput) Percent() float32 {
	return float32(a.read()) / CalibratedMaxAI
}

func (a *AnalogInput) ReadVoltage() float32 {
	return a.Percent() * MaxVoltage
}

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

type Knob struct {
	machine.ADC
	samples int
}

func NewKnob(pin machine.Pin) *Knob {
	adc := machine.ADC{Pin: pin}
	adc.Configure(machine.ADCConfig{})
	return &Knob{ADC: adc, samples: DefaultSamples}
}

func (k *Knob) Samples(samples int) {
	k.samples = samples
}

func (k *Knob) Percent() float32 {
	return 1 - float32(k.read())/math.MaxUint16
}

func (k *Knob) ReadVoltage() float32 {
	return k.Percent() * MaxVoltage
}

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
