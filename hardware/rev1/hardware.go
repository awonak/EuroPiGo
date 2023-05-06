package rev1

import "github.com/awonak/EuroPiGo/hardware/hal"

// aliases for module revision specific referencing
const (
	// DI: Digital Input
	HardwareIdDigital1Input = hal.HardwareIdDigital1Input
	// AI: Analog(ue) Input
	HardwareIdAnalog1Input = hal.HardwareIdAnalog1Input
	// Display
	HardwareIdDisplay1Output = hal.HardwareIdDisplay1Output
	// K1
	HardwareIdKnob1Input = hal.HardwareIdKnob1Input
	// K2
	HardwareIdKnob2Input = hal.HardwareIdKnob2Input
	// B1
	HardwareIdButton1Input = hal.HardwareIdButton1Input
	// B2
	HardwareIdButton2Input = hal.HardwareIdButton2Input
	// CV1
	HardwareIdCV1Output = hal.HardwareIdVoltage1Output
	// CV2
	HardwareIdCV2Output = hal.HardwareIdVoltage2Output
	// CV3
	HardwareIdCV3Output = hal.HardwareIdVoltage3Output
	// CV4
	HardwareIdCV4Output = hal.HardwareIdVoltage4Output
	// CV5
	HardwareIdCV5Output = hal.HardwareIdVoltage5Output
	// CV6
	HardwareIdCV6Output = hal.HardwareIdVoltage6Output
	// RNG
	HardwareIdRandom1Generator = hal.HardwareIdRandom1Generator
)

// aliases for friendly internationali(s|z)ation, colloquialisms, and naming conventions
const (
	HardwareIdAnalogue1Input = HardwareIdAnalog1Input
	HardwareIdOLED1Output    = HardwareIdDisplay1Output
)
