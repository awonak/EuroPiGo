package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

func DoInit() {
	rev1.Initialize(rev1.InitializationParameters{
		InputDigital1:          newNonPicoDigitalReader(bus, hal.HardwareIdDigital1Input),
		InputAnalog1:           newNonPicoAdc(bus, hal.HardwareIdAnalog1Input),
		OutputDisplay1:         newNonPicoDisplayOutput(bus, hal.HardwareIdDisplay1Output),
		InputButton1:           newNonPicoDigitalReader(bus, hal.HardwareIdButton1Input),
		InputButton2:           newNonPicoDigitalReader(bus, hal.HardwareIdButton2Input),
		InputKnob1:             newNonPicoAdc(bus, hal.HardwareIdKnob1Input),
		InputKnob2:             newNonPicoAdc(bus, hal.HardwareIdKnob2Input),
		OutputVoltage1:         newNonPicoPwm(bus, hal.HardwareIdVoltage1Output),
		OutputVoltage2:         newNonPicoPwm(bus, hal.HardwareIdVoltage2Output),
		OutputVoltage3:         newNonPicoPwm(bus, hal.HardwareIdVoltage3Output),
		OutputVoltage4:         newNonPicoPwm(bus, hal.HardwareIdVoltage4Output),
		OutputVoltage5:         newNonPicoPwm(bus, hal.HardwareIdVoltage5Output),
		OutputVoltage6:         newNonPicoPwm(bus, hal.HardwareIdVoltage6Output),
		DeviceRandomGenerator1: nil,
	})
}
