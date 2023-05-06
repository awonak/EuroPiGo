package rev0

import (
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
)

func DoInit() {
	ajCalMap := envelope.NewLerpMap32(rev0.VoltageOutputCalibrationPoints)
	djCalMap := envelope.NewPointMap32(rev0.VoltageOutputCalibrationPoints)
	rev0.Initialize(rev0.InitializationParameters{
		InputButton1:           common.NewNonPicoDigitalReader(bus, rev0.HardwareIdButton1Input),
		InputButton2:           common.NewNonPicoDigitalReader(bus, rev0.HardwareIdButton2Input),
		InputKnob1:             common.NewNonPicoAdc(bus, rev0.HardwareIdKnob1Input),
		InputKnob2:             common.NewNonPicoAdc(bus, rev0.HardwareIdKnob2Input),
		OutputAnalog1:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog1Output, ajCalMap),
		OutputAnalog2:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog2Output, ajCalMap),
		OutputAnalog3:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog3Output, ajCalMap),
		OutputAnalog4:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog4Output, ajCalMap),
		OutputDigital1:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital1Output, djCalMap),
		OutputDigital2:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital2Output, djCalMap),
		OutputDigital3:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital3Output, djCalMap),
		OutputDigital4:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital4Output, djCalMap),
		DeviceRandomGenerator1: nil,
	})
}
