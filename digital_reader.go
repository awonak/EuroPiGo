package europi

import (
	"machine"
	"time"
)

const DefaultDebounceDelay = time.Duration(50 * time.Millisecond)

// DigitalReader is an interface for common digital inputs methods.
type DigitalReader interface {
	Handler(func(machine.Pin))
	HandlerWithDebounce(func(machine.Pin), time.Duration)
	LastInput() time.Time
	Value() bool
}

// DigitalInput is a struct for handling reading of the digital input.
type DigitalInput struct {
	Pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(p machine.Pin)
}

// NewDI creates a new DigitalInput struct.
func NewDI(pin machine.Pin) *DigitalInput {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &DigitalInput{
		Pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// LastInput return the time of the last high input (triggered at 0.8v).
func (d *DigitalInput) LastInput() time.Time {
	return d.lastInput
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *DigitalInput) Value() bool {
	// Invert signal to match expected behavior.
	return !d.Pin.Get()
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *DigitalInput) Handler(handler func(p machine.Pin)) {
	d.HandlerWithDebounce(handler, 0)
}

// Handler sets the callback function to be call when a rising edge is detected and debounce delay time has elapsed.
func (d *DigitalInput) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	d.callback = handler
	d.debounceDelay = delay
	d.Pin.SetInterrupt(machine.PinFalling, d.debounceWrapper)
}

func (d *DigitalInput) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(d.lastInput.Add(d.debounceDelay)) {
		return
	}
	d.callback(p)
	d.lastInput = t
}

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
