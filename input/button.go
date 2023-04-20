package input

import (
	"machine"
	"runtime/interrupt"
	"time"
)

// Button is a struct for handling push button behavior.
type Button struct {
	Pin        machine.Pin
	lastChange time.Time
}

// NewButton creates a new Button struct.
func NewButton(pin machine.Pin) *Button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Button{
		Pin:        pin,
		lastChange: time.Now(),
	}
}

// Handler sets the callback function to be call when the button is pressed.
func (b *Button) Handler(handler func(p machine.Pin)) {
	if handler == nil {
		panic("cannot set nil handler")
	}
	b.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		if b.Value() {
			handler(p)
		}
	})
}

// HandlerEx sets the callback function to be call when the button changes in a specified way.
func (b *Button) HandlerEx(pinChange machine.PinChange, handler func(p machine.Pin)) {
	if handler == nil {
		panic("cannot set nil handler")
	}
	b.setHandler(pinChange, handler)
}

// HandlerWithDebounce sets the callback function to be call when the button is pressed and debounce delay time has elapsed.
func (b *Button) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	if handler == nil {
		panic("cannot set nil handler")
	}
	lastInput := time.Now()
	b.Handler(func(p machine.Pin) {
		now := time.Now()
		if now.Before(lastInput.Add(delay)) {
			return
		}
		handler(p)
		lastInput = now
	})
}

func (b *Button) setHandler(pinChange machine.PinChange, handler func(p machine.Pin)) {
	state := interrupt.Disable()
	b.Pin.SetInterrupt(pinChange, func(p machine.Pin) {
		now := time.Now()
		handler(p)
		b.lastChange = now
	})
	interrupt.Restore(state)
}

// LastChange return the time of the last button input change.
func (b *Button) LastChange() time.Time {
	return b.lastChange
}

// Value returns true if button is currently pressed, else false.
func (b *Button) Value() bool {
	state := interrupt.Disable()
	// Invert signal to match expected behavior.
	v := !b.Pin.Get()
	interrupt.Restore(state)
	return v
}
