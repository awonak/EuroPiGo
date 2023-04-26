package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// These will be configured during `init()` from platform-specific files.
// See `hardware/pico/pico.go` and `hardware/nonpico/nonpico.go` for more information.
var (
	InputDigital1          hal.DigitalInput
	InputAnalog1           hal.AnalogInput
	OutputDisplay1         hal.DisplayOutput
	InputButton1           hal.ButtonInput
	InputButton2           hal.ButtonInput
	InputKnob1             hal.KnobInput
	InputKnob2             hal.KnobInput
	OutputVoltage1         hal.VoltageOutput
	OutputVoltage2         hal.VoltageOutput
	OutputVoltage3         hal.VoltageOutput
	OutputVoltage4         hal.VoltageOutput
	OutputVoltage5         hal.VoltageOutput
	OutputVoltage6         hal.VoltageOutput
	DeviceRandomGenerator1 hal.RandomGenerator
)

// GetHardware returns a EuroPi hardware device based on hardware `id`.
// a `nil` result means that the hardware was not found or some sort of error occurred.
func GetHardware[T any](hw hal.HardwareId) T {
	switch hw {
	case hal.HardwareIdDigital1Input:
		t, _ := InputDigital1.(T)
		return t
	case hal.HardwareIdAnalog1Input:
		t, _ := InputAnalog1.(T)
		return t
	case hal.HardwareIdDisplay1Output:
		t, _ := OutputDisplay1.(T)
		return t
	case hal.HardwareIdButton1Input:
		t, _ := InputButton1.(T)
		return t
	case hal.HardwareIdButton2Input:
		t, _ := InputButton2.(T)
		return t
	case hal.HardwareIdKnob1Input:
		t, _ := InputKnob1.(T)
		return t
	case hal.HardwareIdKnob2Input:
		t, _ := InputKnob2.(T)
		return t
	case hal.HardwareIdVoltage1Output:
		t, _ := OutputVoltage1.(T)
		return t
	case hal.HardwareIdVoltage2Output:
		t, _ := OutputVoltage2.(T)
		return t
	case hal.HardwareIdVoltage3Output:
		t, _ := OutputVoltage3.(T)
		return t
	case hal.HardwareIdVoltage4Output:
		t, _ := OutputVoltage4.(T)
		return t
	case hal.HardwareIdVoltage5Output:
		t, _ := OutputVoltage5.(T)
		return t
	case hal.HardwareIdVoltage6Output:
		t, _ := OutputVoltage6.(T)
		return t
	case hal.HardwareIdRandom1Generator:
		t, _ := DeviceRandomGenerator1.(T)
		return t
	default:
		var none T
		return none
	}
}

// Initialize sets up the hardware
//
// This is only to be called by the automatic platform initialization functions
func Initialize(params InitializationParameters) {
	InputDigital1 = newDigitalInput(params.InputDigital1)
	InputAnalog1 = newAnalogInput(params.InputAnalog1)
	OutputDisplay1 = newDisplayOutput(params.OutputDisplay1)
	InputButton1 = newDigitalInput(params.InputButton1)
	InputButton2 = newDigitalInput(params.InputButton2)
	InputKnob1 = newAnalogInput(params.InputKnob1)
	InputKnob2 = newAnalogInput(params.InputKnob2)
	OutputVoltage1 = newVoltageOuput(params.OutputVoltage1)
	OutputVoltage2 = newVoltageOuput(params.OutputVoltage2)
	OutputVoltage3 = newVoltageOuput(params.OutputVoltage3)
	OutputVoltage4 = newVoltageOuput(params.OutputVoltage4)
	OutputVoltage5 = newVoltageOuput(params.OutputVoltage5)
	OutputVoltage6 = newVoltageOuput(params.OutputVoltage6)
	DeviceRandomGenerator1 = newRandomGenerator(params.DeviceRandomGenerator1)
}

// InitializationParameters is a ferry for hardware functions to the interface layer found here
//
// This is only to be used by the automatic platform initialization functions
type InitializationParameters struct {
	InputDigital1          DigitalReaderProvider
	InputAnalog1           ADCProvider
	OutputDisplay1         DisplayProvider
	InputButton1           DigitalReaderProvider
	InputButton2           DigitalReaderProvider
	InputKnob1             ADCProvider
	InputKnob2             ADCProvider
	OutputVoltage1         PWMProvider
	OutputVoltage2         PWMProvider
	OutputVoltage3         PWMProvider
	OutputVoltage4         PWMProvider
	OutputVoltage5         PWMProvider
	OutputVoltage6         PWMProvider
	DeviceRandomGenerator1 RNDProvider
}
