package europi

import (
	"machine"
	"time"
)

const (
	DefaultDebounceDelay = time.Duration(50 * time.Millisecond)

	DIPin = machine.GPIO22
	B1Pin = machine.GPIO4
	B2Pin = machine.GPIO5
)

var (
	DI DigitalReader
	B1 DigitalReader
	B2 DigitalReader
)

func init() {
	DI = newDI(DIPin)
	B1 = newButton(B1Pin)
	B2 = newButton(B2Pin)
}

// DigitalReader is an interface for common digital inputs methods.
type DigitalReader interface {
	Handler(func(machine.Pin))
	HandlerWithDebounce(func(machine.Pin), time.Duration)
	LastInput() time.Time
	Value() bool
}

// digitalInput is a struct for handling reading of the digital input.
type digitalInput struct {
	Pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(p machine.Pin)
}

// NewDI creates a new DigitalInput struct.
func newDI(pin machine.Pin) *digitalInput {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &digitalInput{
		Pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// LastInput return the time of the last high input (triggered at 0.8v).
func (d *digitalInput) LastInput() time.Time {
	return d.lastInput
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *digitalInput) Value() bool {
	// Invert signal to match expected behavior.
	return !d.Pin.Get()
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *digitalInput) Handler(handler func(p machine.Pin)) {
	d.HandlerWithDebounce(handler, 0)
}

// Handler sets the callback function to be call when a rising edge is detected and debounce delay time has elapsed.
func (d *digitalInput) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	d.callback = handler
	d.debounceDelay = delay
	d.Pin.SetInterrupt(machine.PinFalling, d.debounceWrapper)
}

func (d *digitalInput) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(d.lastInput.Add(d.debounceDelay)) {
		return
	}
	d.callback(p)
	d.lastInput = t
}

// button is a struct for handling push button behavior.
type button struct {
	Pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(p machine.Pin)
}

// newButton creates a new Button struct.
func newButton(pin machine.Pin) *button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &button{
		Pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// Handler sets the callback function to be call when the button is pressed.
func (b *button) Handler(handler func(p machine.Pin)) {
	b.HandlerWithDebounce(handler, 0)
}

// Handler sets the callback function to be call when the button is pressed and debounce delay time has elapsed.
func (b *button) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	b.callback = handler
	b.debounceDelay = delay
	b.Pin.SetInterrupt(machine.PinFalling, b.debounceWrapper)
}

func (b *button) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(b.lastInput.Add(b.debounceDelay)) {
		return
	}
	b.callback(p)
	b.lastInput = t
}

// LastInput return the time of the last button press.
func (b *button) LastInput() time.Time {
	return b.lastInput
}

// Value returns true if button is currently pressed, else false.
func (b *button) Value() bool {
	// Invert signal to match expected behavior.
	return !b.Pin.Get()
}
