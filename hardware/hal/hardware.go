package hal

// HardwareId defines an identifier for specific hardware. See the README.md in the hardware directory for more details.
type HardwareId int

const (
	HardwareIdInvalid = HardwareId(iota)
	HardwareIdRevisionMarker
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
	HardwareIdVoltage7Output
	HardwareIdVoltage8Output
	// NOTE: always ONLY append to this list, NEVER remove, rename, or reorder
)

// aliases for friendly internationali(s|z)ation, colloquialisms, and naming conventions
const (
	HardwareIdAnalogue1Input = HardwareIdAnalog1Input
)

// Hardware is the collection of component wrappers used to interact with the module.
type Hardware interface {
	Revision() Revision
	Random() RandomGenerator
	Button(idx int) ButtonInput
	Knob(idx int) KnobInput
}
