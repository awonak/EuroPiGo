package rev1

import (
	"github.com/heucuva/europi/internal/hardware/hal"
)

var (
	hwDigital1Input    hal.DigitalInput
	hwAnalog1Input     hal.AnalogInput
	hwDisplay1Output   hal.DisplayOutput
	hwButton1Input     hal.ButtonInput
	hwButton2Input     hal.ButtonInput
	hwKnob1Input       hal.KnobInput
	hwKnob2Input       hal.KnobInput
	hwCV1Output        hal.VoltageOutput
	hwCV2Output        hal.VoltageOutput
	hwCV3Output        hal.VoltageOutput
	hwCV4Output        hal.VoltageOutput
	hwCV5Output        hal.VoltageOutput
	hwCV6Output        hal.VoltageOutput
	hwRandom1Generator hal.RandomGenerator
)

func GetHardware(hw hal.HardwareId) any {
	switch hw {
	case hal.HardwareIdDigital1Input:
		return hwDigital1Input
	case hal.HardwareIdAnalog1Input:
		return hwAnalog1Input
	case hal.HardwareIdDisplay1Output:
		return hwDisplay1Output
	case hal.HardwareIdButton1Input:
		return hwButton1Input
	case hal.HardwareIdButton2Input:
		return hwButton2Input
	case hal.HardwareIdKnob1Input:
		return hwKnob1Input
	case hal.HardwareIdKnob2Input:
		return hwKnob2Input
	case hal.HardwareIdVoltage1Output:
		return hwCV1Output
	case hal.HardwareIdVoltage2Output:
		return hwCV2Output
	case hal.HardwareIdVoltage3Output:
		return hwCV3Output
	case hal.HardwareIdVoltage4Output:
		return hwCV4Output
	case hal.HardwareIdVoltage5Output:
		return hwCV5Output
	case hal.HardwareIdVoltage6Output:
		return hwCV6Output
	case hal.HardwareIdRandom1Generator:
		return hwRandom1Generator
	default:
		return nil
	}
}
