//go:build !pico
// +build !pico

package common

import (
	"fmt"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/common"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type nonPicoAdc struct {
	id    hal.HardwareId
	value uint16
}

var (
	// static check
	_ common.ADCProvider = (*nonPicoAdc)(nil)
)

func NewNonPicoAdc(id hal.HardwareId) *nonPicoAdc {
	adc := &nonPicoAdc{
		id: id,
	}
	event.Subscribe(bus, fmt.Sprintf("hw_value_%d", id), func(msg HwMessageADCValue) {
		adc.value = msg.Value
	})
	return adc
}

func (a *nonPicoAdc) Get(samples int) uint16 {
	var sum int
	for i := 0; i < samples; i++ {
		sum += int(a.value)
	}
	return uint16(sum / samples)
}
