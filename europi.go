package europi // import europi "github.com/awonak/EuroPiGo"

import (
	"machine"
)

const (
	MaxVoltage = 10.0
	MinVoltage = 0.0
)

var e *EuroPi

// EuroPi is the collection of component wrappers used to interact with the module.
type EuroPi struct {
	// Display is a wrapper around ssd1306.Device
	Display *display

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
	CV  [6]Outputer
}

func init() {
	if e == nil {
		e = newEuroPi()
	}
}

func New() *EuroPi {
	return e
}

// newEuroPi will return a new EuroPi struct.
func newEuroPi() *EuroPi {
	cv1 := newOutput(machine.GPIO21, machine.PWM2)
	cv2 := newOutput(machine.GPIO20, machine.PWM2)
	cv3 := newOutput(machine.GPIO16, machine.PWM0)
	cv4 := newOutput(machine.GPIO17, machine.PWM0)
	cv5 := newOutput(machine.GPIO18, machine.PWM1)
	cv6 := newOutput(machine.GPIO19, machine.PWM1)

	return &EuroPi{
		Display: newDisplay(machine.I2C0, machine.GPIO0, machine.GPIO1),

		DI: newDI(machine.GPIO22),
		AI: newAI(machine.ADC0),

		B1: newButton(machine.GPIO4),
		B2: newButton(machine.GPIO5),

		K1: newKnob(machine.ADC1),
		K2: newKnob(machine.ADC2),

		CV1: cv1,
		CV2: cv2,
		CV3: cv3,
		CV4: cv4,
		CV5: cv5,
		CV6: cv5,
		CV:  [6]Outputer{cv1, cv2, cv3, cv4, cv5, cv6},
	}
}
