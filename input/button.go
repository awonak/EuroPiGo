package input

import (
	"machine"
	"time"
)

// Button is a struct for handling push button behavior.
type Button struct {
	Pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(p machine.Pin)
}

// NewButton creates a new Button struct.
func NewButton(pin machine.Pin) *Button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Button{
		Pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// Handler sets the callback function to be call when the button is pressed.
func (b *Button) Handler(handler func(p machine.Pin)) {
	b.HandlerWithDebounce(handler, 0)
}

// Handler sets the callback function to be call when the button is pressed and debounce delay time has elapsed.
func (b *Button) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	b.callback = handler
	b.debounceDelay = delay
	b.Pin.SetInterrupt(machine.PinFalling, b.debounceWrapper)
}

func (b *Button) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(b.lastInput.Add(b.debounceDelay)) {
		return
	}
	b.callback(p)
	b.lastInput = t
}

// LastInput return the time of the last button press.
func (b *Button) LastInput() time.Time {
	return b.lastInput
}

// Value returns true if button is currently pressed, else false.
func (b *Button) Value() bool {
	// Invert signal to match expected behavior.
	return !b.Pin.Get()
}
