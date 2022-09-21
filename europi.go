package europi // import europi "github.com/awonak/EuroPiGo"

import (
	"machine"
)

const (
	MaxVoltage = 10.0
	MinVoltage = 0.0
)

// EuroPi is the collection of component wrappers used to interact with the module.
type EuroPi struct {
	// Display is a wrapper around ssd1306.Device
	Display *Display

	DI DigitalReader
	AI AnalogReader

	B1 DigitalReader
	B2 DigitalReader

	K1 AnalogReader
	K2 AnalogReader

	CV1 Outputer
	CV2 Outputer
	CV3 Outputer
	CV4 Outputer
	CV5 Outputer
	CV6 Outputer
	CV  []Outputer
}

// New will return a new EuroPi struct.
func New() *EuroPi {
	cv1 := NewOutput(machine.GPIO21, machine.PWM2)
	cv2 := NewOutput(machine.GPIO20, machine.PWM2)
	cv3 := NewOutput(machine.GPIO16, machine.PWM0)
	cv4 := NewOutput(machine.GPIO17, machine.PWM0)
	cv5 := NewOutput(machine.GPIO18, machine.PWM1)
	cv6 := NewOutput(machine.GPIO19, machine.PWM1)

	return &EuroPi{
		Display: NewDisplay(machine.I2C0, machine.GPIO0, machine.GPIO1),

		DI: NewDI(machine.GPIO22),
		AI: NewAI(machine.ADC0),

		B1: NewButton(machine.GPIO4),
		B2: NewButton(machine.GPIO5),

		K1: NewKnob(machine.ADC1),
		K2: NewKnob(machine.ADC2),

		CV1: cv1,
		CV2: cv2,
		CV3: cv3,
		CV4: cv4,
		CV5: cv5,
		CV6: cv5,
		CV:  []Outputer{cv1, cv2, cv3, cv4, cv5, cv6},
	}
}
