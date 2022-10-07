package europi // import europi "github.com/awonak/EuroPiGo"

import (
	"machine"
)

const (
	MaxVoltage = 10.0
	MinVoltage = 0.0

	// EuroPi hardware GPIO pins
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
)
