package rev0

import (
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
)

func DoInit() {
	rev0.Initialize(rev0.InitializationParameters{
		InputButton1:           common.NewNonPicoDigitalReader(rev0.HardwareIdButton1Input),
		InputButton2:           common.NewNonPicoDigitalReader(rev0.HardwareIdButton2Input),
		InputKnob1:             common.NewNonPicoAdc(rev0.HardwareIdKnob1Input),
		InputKnob2:             common.NewNonPicoAdc(rev0.HardwareIdKnob2Input),
		OutputAnalog1:          common.NewNonPicoPwm(rev0.HardwareIdAnalog1Output, rev0.DefaultVoltageOutputCalibration),
		OutputAnalog2:          common.NewNonPicoPwm(rev0.HardwareIdAnalog2Output, rev0.DefaultVoltageOutputCalibration),
		OutputAnalog3:          common.NewNonPicoPwm(rev0.HardwareIdAnalog3Output, rev0.DefaultVoltageOutputCalibration),
		OutputAnalog4:          common.NewNonPicoPwm(rev0.HardwareIdAnalog4Output, rev0.DefaultVoltageOutputCalibration),
		OutputDigital1:         common.NewNonPicoPwm(rev0.HardwareIdDigital1Output, rev0.DefaultVoltageOutputCalibration),
		OutputDigital2:         common.NewNonPicoPwm(rev0.HardwareIdDigital2Output, rev0.DefaultVoltageOutputCalibration),
		OutputDigital3:         common.NewNonPicoPwm(rev0.HardwareIdDigital3Output, rev0.DefaultVoltageOutputCalibration),
		OutputDigital4:         common.NewNonPicoPwm(rev0.HardwareIdDigital4Output, rev0.DefaultVoltageOutputCalibration),
		DeviceRandomGenerator1: nil,
	})
}
