package rev1

import (
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
)

func DoInit() {
	cvCalMap := envelope.NewLerpMap32(rev1.VoltageOutputCalibrationPoints)
	rev1.Initialize(rev1.InitializationParameters{
		InputDigital1:          common.NewNonPicoDigitalReader(bus, rev1.HardwareIdDigital1Input),
		InputAnalog1:           common.NewNonPicoAdc(bus, rev1.HardwareIdAnalog1Input),
		OutputDisplay1:         common.NewNonPicoDisplayOutput(bus, rev1.HardwareIdDisplay1Output),
		InputButton1:           common.NewNonPicoDigitalReader(bus, rev1.HardwareIdButton1Input),
		InputButton2:           common.NewNonPicoDigitalReader(bus, rev1.HardwareIdButton2Input),
		InputKnob1:             common.NewNonPicoAdc(bus, rev1.HardwareIdKnob1Input),
		InputKnob2:             common.NewNonPicoAdc(bus, rev1.HardwareIdKnob2Input),
		OutputVoltage1:         common.NewNonPicoPwm(bus, rev1.HardwareIdCV1Output, cvCalMap),
		OutputVoltage2:         common.NewNonPicoPwm(bus, rev1.HardwareIdCV2Output, cvCalMap),
		OutputVoltage3:         common.NewNonPicoPwm(bus, rev1.HardwareIdCV3Output, cvCalMap),
		OutputVoltage4:         common.NewNonPicoPwm(bus, rev1.HardwareIdCV4Output, cvCalMap),
		OutputVoltage5:         common.NewNonPicoPwm(bus, rev1.HardwareIdCV5Output, cvCalMap),
		OutputVoltage6:         common.NewNonPicoPwm(bus, rev1.HardwareIdCV6Output, cvCalMap),
		DeviceRandomGenerator1: nil,
	})
}
