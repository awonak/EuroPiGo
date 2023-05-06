//go:build !pico
// +build !pico

package common

import (
	"fmt"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

var (
	bus = event.NewBus()
)

func SetDigitalValue(hid hal.HardwareId, value bool) {
	bus.Post(fmt.Sprintf("hw_value_%d", hid), HwMessageDigitalValue{
		Value: value,
	})
}

func TriggerInterrupt(hid hal.HardwareId, change hal.ChangeFlags) {
	bus.Post(fmt.Sprintf("hw_interrupt_%d", hid), HwMessageInterrupt{
		Change: change,
	})
}

func SetADCValue(hid hal.HardwareId, value uint16) {
	bus.Post(fmt.Sprintf("hw_value_%d", hid), HwMessageADCValue{
		Value: value,
	})
}

func OnPWMValue(hid hal.HardwareId, fn func(hid hal.HardwareId, value uint16, voltage float32)) {
	event.Subscribe(bus, fmt.Sprintf("hw_pwm_%d", hid), func(msg HwMessagePwmValue) {
		fn(hid, msg.Value, msg.Voltage)
	})
}

func OnDisplayOutput(hid hal.HardwareId, fn func(hid hal.HardwareId, op HwDisplayOp, params []int16)) {
	event.Subscribe(bus, fmt.Sprintf("hw_display_%d", hid), func(msg HwMessageDisplay) {
		fn(hid, msg.Op, msg.Operands)
	})
}
