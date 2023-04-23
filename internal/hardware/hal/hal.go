package hal

type HardwareId int

const (
	HardwareIdInvalid = HardwareId(iota)
	HardwareIdDigital1Input
	HardwareIdAnalog1Input
	HardwareIdDisplay1Output
	HardwareIdButton1Input
	HardwareIdButton2Input
	HardwareIdKnob1Input
	HardwareIdKnob2Input
	HardwareIdVoltage1Output
	HardwareIdVoltage2Output
	HardwareIdVoltage3Output
	HardwareIdVoltage4Output
	HardwareIdVoltage5Output
	HardwareIdVoltage6Output
	HardwareIdRandom1Generator
	// NOTE: always ONLY append to this list, NEVER remove, rename, or reorder
)

// aliases for friendly internationali(s|z)ation
const (
	HardwareIdAnalogue1Input = HardwareIdAnalog1Input
)
