//go:build !pico
// +build !pico

package rev0

import (
	"sync"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
	"github.com/awonak/EuroPiGo/lerp"
)

func setupDefaultState() {
	common.SetDigitalValue(rev0.HardwareIdButton1Input, false)
	common.SetDigitalValue(rev0.HardwareIdButton2Input, false)

	common.SetADCValue(rev0.HardwareIdKnob1Input, aiLerp.Lerp(0.5))
	common.SetADCValue(rev0.HardwareIdKnob2Input, aiLerp.Lerp(0.5))
}

func setupVoltageOutputListeners(cb func(id hal.HardwareId, voltage float32)) {
	ids := []hal.HardwareId{
		rev0.HardwareIdAnalog1Output,
		rev0.HardwareIdAnalog2Output,
		rev0.HardwareIdAnalog3Output,
		rev0.HardwareIdAnalog4Output,
		rev0.HardwareIdDigital1Output,
		rev0.HardwareIdDigital2Output,
		rev0.HardwareIdDigital3Output,
		rev0.HardwareIdDigital4Output,
	}
	for _, id := range ids {
		common.OnPWMValue(id, func(hid hal.HardwareId, value uint16, voltage float32) {
			cb(hid, voltage)
		})
	}
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
	aiLerp = lerp.NewLerp32[uint16](rev0.DefaultCalibratedMinAI, rev0.DefaultCalibratedMaxAI)
)

func setAnalogInput(id hal.HardwareId, voltage float32) {
	common.SetADCValue(id, aiLerp.Lerp(voltage))
}
