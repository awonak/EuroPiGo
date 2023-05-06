//go:build !pico
// +build !pico

package rev1

import (
	"sync"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
	"github.com/awonak/EuroPiGo/lerp"
)

func setupDefaultState() {
	common.SetDigitalValue(rev1.HardwareIdDigital1Input, false)
	common.SetADCValue(rev1.HardwareIdAnalog1Input, rev1.DefaultCalibratedMaxAI)

	common.SetDigitalValue(rev1.HardwareIdButton1Input, false)
	common.SetDigitalValue(rev1.HardwareIdButton2Input, false)

	common.SetADCValue(rev1.HardwareIdKnob1Input, aiLerp.Lerp(0.5))
	common.SetADCValue(rev1.HardwareIdKnob2Input, aiLerp.Lerp(0.5))
}

func setupVoltageOutputListeners(cb func(id hal.HardwareId, voltage float32)) {
	for id := hal.HardwareIdVoltage1Output; id <= hal.HardwareIdVoltage6Output; id++ {
		common.OnPWMValue(id, func(hid hal.HardwareId, value uint16, voltage float32) {
			cb(hid, voltage)
		})
	}
}

func setupDisplayOutputListener(cb func(id hal.HardwareId, op common.HwDisplayOp, params []int16)) {
	common.OnDisplayOutput(hal.HardwareIdDisplay1Output, cb)
}

var (
	states sync.Map
)

func setDigitalInput(id hal.HardwareId, value bool) {
	prevState, _ := states.Load(id)

	states.Store(id, value)
	common.SetDigitalValue(id, value)

	if prevState != value {
		if value {
			// rising
			common.TriggerInterrupt(id, hal.ChangeRising)
		} else {
			// falling
			common.TriggerInterrupt(id, hal.ChangeFalling)
		}
	}
}

var (
	aiLerp = lerp.NewLerp32[uint16](rev1.DefaultCalibratedMinAI, rev1.DefaultCalibratedMaxAI)
)

func setAnalogInput(id hal.HardwareId, voltage float32) {
	common.SetADCValue(id, aiLerp.Lerp(voltage))
}
