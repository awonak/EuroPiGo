package europi // import europi "github.com/awonak/EuroPiGo"

import (
	"machine"
)

const (
	// EuroPi voltage range constants.
	MaxVoltage = 10.0
	MinVoltage = 0.0

	// EuroPi hardware GPIO pins.
	DIPin = machine.GPIO22
	AIPin = machine.ADC0

	DisplaySdaPin = machine.GPIO0
	DisplaySclPin = machine.GPIO1

	K1Pin = machine.ADC1
	K2Pin = machine.ADC2

	B1Pin = machine.GPIO4
	B2Pin = machine.GPIO5

	CV1Pin = machine.GPIO21
	CV2Pin = machine.GPIO20
	CV3Pin = machine.GPIO16
	CV4Pin = machine.GPIO17
	CV5Pin = machine.GPIO18
	CV6Pin = machine.GPIO19
)

var (
	DisplayChannel = machine.I2C0
	CV1PwmGroup    = machine.PWM2
	CV2PwmGroup    = machine.PWM2
	CV3PwmGroup    = machine.PWM0
	CV4PwmGroup    = machine.PWM0
	CV5PwmGroup    = machine.PWM1
	CV6PwmGroup    = machine.PWM1

	europi *EuroPi
)

// EuroPi is the collection of component wrappers used to interact with the module.
type EuroPi struct {
	// Display provides methods for drawing to the OLED display.
	Display *display

	DI *digitalInput
	// AI provides methods for reading analog input control voltage between 0 and 12V.
	AI *analogInput

	// B1 is a struct for handling the left push button behavior.
	B1 *button
	// B2 is a struct for handling the right push button behavior.
	B2 *button

	// K1 provides methods for reading knob voltage and position for the left knob.
	K1 *knob
	// K2 provides methods for reading knob voltage and position for the left knob.
	K2 *knob

	// CV1-6 are structs for interacting with the cv output jacks.
	CV1 *output
	CV2 *output
	CV3 *output
	CV4 *output
	CV5 *output
	CV6 *output
	// CV is an array containing all CV outputs.
	CV [6]*output
}

func init() {
	europi = new()
}

func new() *EuroPi {
	cv1 := newOutput(CV1Pin, CV1PwmGroup)
	cv2 := newOutput(CV2Pin, CV2PwmGroup)
	cv3 := newOutput(CV3Pin, CV3PwmGroup)
	cv4 := newOutput(CV4Pin, CV4PwmGroup)
	cv5 := newOutput(CV5Pin, CV5PwmGroup)
	cv6 := newOutput(CV6Pin, CV6PwmGroup)

	return &EuroPi{
		Display: newDisplay(DisplayChannel, DisplaySdaPin, DisplaySclPin),

		DI: newDigitalInput(DIPin),
		AI: newAnalogInput(AIPin),

		B1: newButton(B1Pin),
		B2: newButton(B2Pin),

		K1: newKnob(K1Pin),
		K2: newKnob(K2Pin),

		CV1: cv1,
		CV2: cv2,
		CV3: cv3,
		CV4: cv4,
		CV5: cv5,
		CV6: cv5,
		CV:  [6]*output{cv1, cv2, cv3, cv4, cv5, cv6},
	}
}

func GetInstance() *EuroPi {
	return europi
}
