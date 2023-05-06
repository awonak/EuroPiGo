package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
)

func DoInit() {
	rev1.Initialize(rev1.InitializationParameters{
		InputDigital1:          common.NewNonPicoDigitalReader(rev1.HardwareIdDigital1Input),
		InputAnalog1:           common.NewNonPicoAdc(rev1.HardwareIdAnalog1Input),
		OutputDisplay1:         common.NewNonPicoDisplayOutput(rev1.HardwareIdDisplay1Output),
		InputButton1:           common.NewNonPicoDigitalReader(rev1.HardwareIdButton1Input),
		InputButton2:           common.NewNonPicoDigitalReader(rev1.HardwareIdButton2Input),
		InputKnob1:             common.NewNonPicoAdc(rev1.HardwareIdKnob1Input),
		InputKnob2:             common.NewNonPicoAdc(rev1.HardwareIdKnob2Input),
		OutputVoltage1:         common.NewNonPicoPwm(rev1.HardwareIdCV1Output, rev1.DefaultVoltageOutputCalibration),
		OutputVoltage2:         common.NewNonPicoPwm(rev1.HardwareIdCV2Output, rev1.DefaultVoltageOutputCalibration),
		OutputVoltage3:         common.NewNonPicoPwm(rev1.HardwareIdCV3Output, rev1.DefaultVoltageOutputCalibration),
		OutputVoltage4:         common.NewNonPicoPwm(rev1.HardwareIdCV4Output, rev1.DefaultVoltageOutputCalibration),
		OutputVoltage5:         common.NewNonPicoPwm(rev1.HardwareIdCV5Output, rev1.DefaultVoltageOutputCalibration),
		OutputVoltage6:         common.NewNonPicoPwm(rev1.HardwareIdCV6Output, rev1.DefaultVoltageOutputCalibration),
		DeviceRandomGenerator1: nil,
	})
}
