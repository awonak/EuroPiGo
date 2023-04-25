//go:build !pico && revision1
// +build !pico,revision1

package events

import (
	"fmt"
	"math"
	"sync"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/lerp"
)

var (
	voLerp = lerp.NewLerp32[uint16](0, math.MaxUint16)
)

func SetupVoltageOutputListeners(cb func(id hal.HardwareId, voltage float32)) {
	bus := rev1.DefaultEventBus

	for id := hal.HardwareIdVoltage1Output; id <= hal.HardwareIdVoltage6Output; id++ {
		fn := func(hid hal.HardwareId) func(rev1.HwMessagePwmValue) {
			return func(msg rev1.HwMessagePwmValue) {
				v := voLerp.ClampedInverseLerp(msg.Value) * rev1.MaxOutputVoltage
				cb(hid, v)
			}
		}(id)
		event.Subscribe(bus, fmt.Sprintf("hw_pwm_%d", id), fn)
	}
}

func SetupDisplayOutputListener(cb func(id hal.HardwareId, op rev1.HwDisplayOp, params []int16)) {
	bus := rev1.DefaultEventBus
	id := hal.HardwareIdDisplay1Output
	event.Subscribe(bus, fmt.Sprintf("hw_display_%d", id), func(msg rev1.HwMessageDisplay) {
		cb(id, msg.Op, msg.Operands)
	})

}

var (
	states sync.Map
)

func SetDigitalInput(id hal.HardwareId, value bool) {
	prevState, _ := states.Load(id)

	bus := rev1.DefaultEventBus

	states.Store(id, value)
	bus.Post(fmt.Sprintf("hw_value_%d", id), rev1.HwMessageDigitalValue{
		Value: value,
	})

	if prevState != value {
		if value {
			// rising
			bus.Post(fmt.Sprintf("hw_interrupt_%d", id), rev1.HwMessageInterrupt{
				Change: hal.ChangeRising,
			})
		} else {
			// falling
			bus.Post(fmt.Sprintf("hw_interrupt_%d", id), rev1.HwMessageInterrupt{
				Change: hal.ChangeFalling,
			})
		}
	}
}

var (
	aiLerp = lerp.NewLerp32[uint16](rev1.DefaultCalibratedMinAI, rev1.DefaultCalibratedMaxAI)
)

func SetAnalogInput(id hal.HardwareId, voltage float32) {
	bus := rev1.DefaultEventBus

	bus.Post(fmt.Sprintf("hw_value_%d", id), rev1.HwMessageADCValue{
		Value: aiLerp.Lerp(voltage),
	})
}
