//go:build pico
// +build pico

package pico

import (
	"machine"
	"runtime/interrupt"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

type picoDigitalReader struct {
	pin machine.Pin
}

func newPicoDigitalReader(pin machine.Pin) *picoDigitalReader {
	dr := &picoDigitalReader{
		pin: pin,
	}
	dr.pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return dr
}

func (d *picoDigitalReader) Get() bool {
	state := interrupt.Disable()
	// Invert signal to match expected behavior.
	v := !d.pin.Get()
	interrupt.Restore(state)
	return v
}

func (d *picoDigitalReader) SetHandler(changes hal.ChangeFlags, handler func()) {
	pinChange := d.convertChangeFlags(changes)

	state := interrupt.Disable()
	d.pin.SetInterrupt(pinChange, func(machine.Pin) {
		handler()
	})
	interrupt.Restore(state)
}

func (d *picoDigitalReader) convertChangeFlags(changes hal.ChangeFlags) machine.PinChange {
	var pinChange machine.PinChange
	if (changes & hal.ChangeRising) != 0 {
		pinChange |= machine.PinFalling
	}
	if (changes & hal.ChangeFalling) != 0 {
		pinChange |= machine.PinRising
	}
	return pinChange
}
