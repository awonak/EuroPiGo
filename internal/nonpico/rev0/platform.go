package rev0

import (
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
)

func DoInit() {
	cvCalMap := envelope.NewMap32([]envelope.MapEntry[float32, uint16]{
		{
			Input:  rev0.MinOutputVoltage,
			Output: rev0.CalibratedTop,
		},
		{
			Input:  rev0.MaxOutputVoltage,
			Output: rev0.CalibratedOffset,
		},
	})
	rev0.Initialize(rev0.InitializationParameters{
		InputButton1:           common.NewNonPicoDigitalReader(bus, rev0.HardwareIdButton1Input),
		InputButton2:           common.NewNonPicoDigitalReader(bus, rev0.HardwareIdButton2Input),
		InputKnob1:             common.NewNonPicoAdc(bus, rev0.HardwareIdKnob1Input),
		InputKnob2:             common.NewNonPicoAdc(bus, rev0.HardwareIdKnob2Input),
		OutputAnalog1:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog1Output, cvCalMap),
		OutputAnalog2:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog2Output, cvCalMap),
		OutputAnalog3:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog3Output, cvCalMap),
		OutputAnalog4:          common.NewNonPicoPwm(bus, rev0.HardwareIdAnalog4Output, cvCalMap),
		OutputDigital1:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital1Output, cvCalMap),
		OutputDigital2:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital2Output, cvCalMap),
		OutputDigital3:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital3Output, cvCalMap),
		OutputDigital4:         common.NewNonPicoPwm(bus, rev0.HardwareIdDigital4Output, cvCalMap),
		DeviceRandomGenerator1: nil,
	})
}
