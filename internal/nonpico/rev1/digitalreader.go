//go:build !pico
// +build !pico

package rev1

import (
	"fmt"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

type nonPicoDigitalReader struct {
	bus   event.Bus
	id    hal.HardwareId
	value bool
}

func newNonPicoDigitalReader(bus event.Bus, id hal.HardwareId) rev1.DigitalReaderProvider {
	dr := &nonPicoDigitalReader{
		bus: bus,
		id:  id,
	}
	event.Subscribe(bus, fmt.Sprintf("hw_value_%d", id), func(msg HwMessageDigitalValue) {
		dr.value = msg.Value
	})
	return dr
}

func (d *nonPicoDigitalReader) Get() bool {
	// Invert signal to match expected behavior.
	return !d.value
}

func (d *nonPicoDigitalReader) SetHandler(changes hal.ChangeFlags, handler func()) {
	event.Subscribe(d.bus, fmt.Sprintf("hw_interrupt_%d", d.id), func(msg HwMessageInterrupt) {
		if (msg.Change & changes) != 0 {
			handler()
		}
	})
}
