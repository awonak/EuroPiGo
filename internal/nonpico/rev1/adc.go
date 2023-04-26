//go:build !pico
// +build !pico

package rev1

import (
	"fmt"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

type nonPicoAdc struct {
	bus   event.Bus
	id    hal.HardwareId
	value uint16
}

func newNonPicoAdc(bus event.Bus, id hal.HardwareId) rev1.ADCProvider {
	adc := &nonPicoAdc{
		bus: bus,
		id:  id,
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
