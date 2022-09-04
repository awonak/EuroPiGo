package europi

import (
	"machine"
)

const (
	MaxVoltage = 10.0
	MinVoltage = 0.0
)

type EuroPi struct {
	DI *DigitalInput
	AI *AnalogInput

	Display *Display

	B1 *Button
	B2 *Button

	K1 *Knob
	K2 *Knob

	CV1 *Output
	CV2 *Output
	CV3 *Output
	CV4 *Output
	CV5 *Output
	CV6 *Output
}

func New() EuroPi {

	europi := EuroPi{
		DI: NewDI(machine.GPIO22),
		AI: NewAI(machine.ADC0),

		Display: NewDisplay(machine.I2C0, machine.GPIO0, machine.GPIO1),

		B1: NewButton(machine.GPIO4),
		B2: NewButton(machine.GPIO5),

		K1: NewKnob(machine.ADC1),
		K2: NewKnob(machine.ADC2),

		CV1: NewOutput(machine.GPIO21, machine.PWM2),
		CV2: NewOutput(machine.GPIO20, machine.PWM2),
		CV3: NewOutput(machine.GPIO16, machine.PWM0),
		CV4: NewOutput(machine.GPIO17, machine.PWM0),
		CV5: NewOutput(machine.GPIO18, machine.PWM1),
		CV6: NewOutput(machine.GPIO19, machine.PWM1),
	}

	return europi
}
