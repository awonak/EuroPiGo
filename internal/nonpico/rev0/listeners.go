//go:build !pico
// +build !pico

package rev0

import (
	"fmt"
	"sync"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/common"
	"github.com/awonak/EuroPiGo/lerp"
)

var (
	bus = event.NewBus()
)

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
		fn := func(hid hal.HardwareId) func(common.HwMessagePwmValue) {
			return func(msg common.HwMessagePwmValue) {
				cb(hid, msg.Voltage)
			}
		}(id)
		event.Subscribe(bus, fmt.Sprintf("hw_pwm_%d", id), fn)
	}
}

var (
	states sync.Map
)

func setDigitalInput(id hal.HardwareId, value bool) {
	prevState, _ := states.Load(id)

	states.Store(id, value)
	bus.Post(fmt.Sprintf("hw_value_%d", id), common.HwMessageDigitalValue{
		Value: value,
	})

	if prevState != value {
		if value {
			// rising
			bus.Post(fmt.Sprintf("hw_interrupt_%d", id), common.HwMessageInterrupt{
				Change: hal.ChangeRising,
			})
		} else {
			// falling
			bus.Post(fmt.Sprintf("hw_interrupt_%d", id), common.HwMessageInterrupt{
				Change: hal.ChangeFalling,
			})
		}
	}
}

var (
	aiLerp = lerp.NewLerp32[uint16](rev0.DefaultCalibratedMinAI, rev0.DefaultCalibratedMaxAI)
)

func setAnalogInput(id hal.HardwareId, voltage float32) {
	bus.Post(fmt.Sprintf("hw_value_%d", id), common.HwMessageADCValue{
		Value: aiLerp.Lerp(voltage),
	})
}
