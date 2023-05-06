//go:build !pico
// +build !pico

package common

import (
	"fmt"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/common"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type nonPicoDigitalReader struct {
	id    hal.HardwareId
	value bool
}

var (
	// static check
	_ common.DigitalReaderProvider = (*nonPicoDigitalReader)(nil)
)

func NewNonPicoDigitalReader(id hal.HardwareId) *nonPicoDigitalReader {
	dr := &nonPicoDigitalReader{
		id:    id,
		value: true, // start off in high, as that's actually read as low
	}
	event.Subscribe(bus, fmt.Sprintf("hw_value_%d", id), func(msg HwMessageDigitalValue) {
		dr.value = !msg.Value
	})
	return dr
}

func (d *nonPicoDigitalReader) Get() bool {
	// Invert signal to match expected behavior.
	return !d.value
}

func (d *nonPicoDigitalReader) SetHandler(changes hal.ChangeFlags, handler func()) {
	event.Subscribe(bus, fmt.Sprintf("hw_interrupt_%d", d.id), func(msg HwMessageInterrupt) {
		if (msg.Change & changes) != 0 {
			handler()
		}
	})
}
