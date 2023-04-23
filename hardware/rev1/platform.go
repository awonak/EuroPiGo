package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

var (
	RevisionMarker         hal.RevisionMarker
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

func GetHardware[T any](hw hal.HardwareId) T {
	switch hw {
	case hal.HardwareIdRevisionMarker:
		t, _ := RevisionMarker.(T)
		return t
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
