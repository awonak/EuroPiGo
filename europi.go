package europi // import "github.com/heucuva/europi"

import (
	"machine"

	"github.com/heucuva/europi/input"
	"github.com/heucuva/europi/output"
)

// EuroPi is the collection of component wrappers used to interact with the module.
type EuroPi struct {
	// Display is a wrapper around ssd1306.Device
	Display *output.Display

	DI input.DigitalReader
	AI input.AnalogReader

	B1 input.DigitalReader
	B2 input.DigitalReader

	K1 input.AnalogReader
	K2 input.AnalogReader

	CV1 output.Output
	CV2 output.Output
	CV3 output.Output
	CV4 output.Output
	CV5 output.Output
	CV6 output.Output
	CV  [6]output.Output
}

// New will return a new EuroPi struct.
func New() *EuroPi {
	cv1 := output.NewOutput(machine.GPIO21, machine.PWM2)
	cv2 := output.NewOutput(machine.GPIO20, machine.PWM2)
	cv3 := output.NewOutput(machine.GPIO16, machine.PWM0)
	cv4 := output.NewOutput(machine.GPIO17, machine.PWM0)
	cv5 := output.NewOutput(machine.GPIO18, machine.PWM1)
	cv6 := output.NewOutput(machine.GPIO19, machine.PWM1)

	e := &EuroPi{
		Display: output.NewDisplay(machine.I2C0, machine.GPIO0, machine.GPIO1),

		DI: input.NewDigital(machine.GPIO22),
		AI: input.NewAnalog(machine.ADC0),

		B1: input.NewButton(machine.GPIO4),
		B2: input.NewButton(machine.GPIO5),

		K1: input.NewKnob(machine.ADC1),
		K2: input.NewKnob(machine.ADC2),

		CV1: cv1,
		CV2: cv2,
		CV3: cv3,
		CV4: cv4,
		CV5: cv5,
		CV6: cv5,
		CV:  [6]output.Output{cv1, cv2, cv3, cv4, cv5, cv6},
	}

	return e
}
