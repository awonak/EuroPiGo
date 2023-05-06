package rev0

import "github.com/awonak/EuroPiGo/hardware/hal"

// aliases for module revision specific referencing
const (
	// K1
	HardwareIdKnob1Input = hal.HardwareIdKnob1Input
	// K2
	HardwareIdKnob2Input = hal.HardwareIdKnob2Input
	// B1
	HardwareIdButton1Input = hal.HardwareIdButton1Input
	// B2
	HardwareIdButton2Input = hal.HardwareIdButton2Input
	// AJ1
	HardwareIdAnalog1Output = hal.HardwareIdVoltage1Output
	// AJ2
	HardwareIdAnalog2Output = hal.HardwareIdVoltage2Output
	// AJ3
	HardwareIdAnalog3Output = hal.HardwareIdVoltage3Output
	// AJ4
	HardwareIdAnalog4Output = hal.HardwareIdVoltage4Output
	// DJ1
	HardwareIdDigital1Output = hal.HardwareIdVoltage5Output
	// DJ2
	HardwareIdDigital2Output = hal.HardwareIdVoltage6Output
	// DJ3
	HardwareIdDigital3Output = hal.HardwareIdVoltage7Output
	// DJ4
	HardwareIdDigital4Output = hal.HardwareIdVoltage8Output
	// RNG
	HardwareIdRandom1Generator = hal.HardwareIdRandom1Generator
)

// aliases for friendly internationali(s|z)ation, colloquialisms, and naming conventions
const (
	HardwareIdAnalogue1Output = HardwareIdAnalog1Output
	HardwareIdAnalogue2Output = HardwareIdAnalog2Output
	HardwareIdAnalogue3Output = HardwareIdAnalog3Output
	HardwareIdAnalogue4Output = HardwareIdAnalog4Output
)
