package input

import (
	"machine"
	"time"
)

const DefaultDebounceDelay = time.Duration(50 * time.Millisecond)

// Digital is a struct for handling reading of the digital input.
type Digital struct {
	Pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(p machine.Pin)
}

// NewDigital creates a new Digital struct.
func NewDigital(pin machine.Pin) *Digital {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Digital{
		Pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// LastInput return the time of the last high input (triggered at 0.8v).
func (d *Digital) LastInput() time.Time {
	return d.lastInput
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *Digital) Value() bool {
	// Invert signal to match expected behavior.
	return !d.Pin.Get()
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *Digital) Handler(handler func(p machine.Pin)) {
	d.HandlerWithDebounce(handler, 0)
}

// Handler sets the callback function to be call when a rising edge is detected and debounce delay time has elapsed.
func (d *Digital) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	d.callback = handler
	d.debounceDelay = delay
	d.Pin.SetInterrupt(machine.PinFalling, d.debounceWrapper)
}

func (d *Digital) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(d.lastInput.Add(d.debounceDelay)) {
		return
	}
	d.callback(p)
	d.lastInput = t
}
