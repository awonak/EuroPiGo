package input

import (
	"machine"
	"runtime/interrupt"
	"time"
)

const DefaultDebounceDelay = time.Duration(50 * time.Millisecond)

// Digital is a struct for handling reading of the digital input.
type Digital struct {
	Pin        machine.Pin
	lastChange time.Time
}

// NewDigital creates a new Digital struct.
func NewDigital(pin machine.Pin) *Digital {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Digital{
		Pin:        pin,
		lastChange: time.Now(),
	}
}

// LastChange return the time of the last input change (triggered at 0.8v).
func (d *Digital) LastChange() time.Time {
	return d.lastChange
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *Digital) Value() bool {
	state := interrupt.Disable()
	// Invert signal to match expected behavior.
	v := !d.Pin.Get()
	interrupt.Restore(state)
	return v
}

// Handler sets the callback function to be call when the falling edge is detected.
func (d *Digital) Handler(handler func(p machine.Pin)) {
	if handler == nil {
		panic("cannot set nil handler")
	}
	d.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		if d.Value() {
			handler(p)
		}
	})
}

// HandlerEx sets the callback function to be call when the input changes in a specified way.
func (d *Digital) HandlerEx(pinChange machine.PinChange, handler func(p machine.Pin)) {
	if handler == nil {
		panic("cannot set nil handler")
	}
	d.setHandler(pinChange, handler)
}

// HandlerWithDebounce sets the callback function to be call when the falling edge is detected and debounce delay time has elapsed.
func (d *Digital) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	if handler == nil {
		panic("cannot set nil handler")
	}
	lastInput := time.Now()
	d.Handler(func(p machine.Pin) {
		now := time.Now()
		if now.Before(lastInput.Add(delay)) {
			return
		}
		handler(p)
		lastInput = now
	})
}

func (d *Digital) setHandler(pinChange machine.PinChange, handler func(p machine.Pin)) {
	state := interrupt.Disable()
	d.Pin.SetInterrupt(pinChange, func(p machine.Pin) {
		now := time.Now()
		handler(p)
		d.lastChange = now
	})
	interrupt.Restore(state)
}
