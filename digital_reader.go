package europi

import (
	"machine"
	"time"
)

const DefaultDebounceDelay = time.Duration(50 * time.Millisecond)

// DigitalReader is an interface for common digital inputs methods.
type DigitalReader interface {
	Handler(func(machine.Pin))
	Debounce(delay time.Duration)
	LastInput() time.Time
	Value() bool
}

// DigitalInput is a struct for handling reading of the digital input.
type DigitalInput struct {
	Pin           machine.Pin
	lastInputTime time.Time
	debounceDelay time.Duration
}

// NewDI creates a new DigitalInput struct.
func NewDI(pin machine.Pin) *DigitalInput {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &DigitalInput{
		Pin:           pin,
		lastInputTime: time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// LastInput return the time of the last high input (triggered at 0.8v).
func (d *DigitalInput) LastInput() time.Time {
	return d.lastInputTime
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *DigitalInput) Value() bool {
	// Invert signal to match expected behavior.
	return !d.Pin.Get()
}

// Debounce overrides the default debounce delay with the provided duration value.
func (d *DigitalInput) Debounce(delay time.Duration) {
	d.debounceDelay = delay
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *DigitalInput) Handler(handler func(p machine.Pin)) {
	d.debounceWrapper(handler)
}
func (d *DigitalInput) debounceWrapper(handler func(p machine.Pin)) {
	wrapped := func(p machine.Pin) {
		if time.Now().Before(d.lastInputTime.Add(d.debounceDelay)) {
			return
		}
		d.lastInputTime = time.Now()
		handler(p)
	}
	d.Pin.SetInterrupt(machine.PinRising, wrapped)
}

// Button is a struct for handling push button behavior.
type Button struct {
	Pin           machine.Pin
	lastInputTime time.Time
	debounceDelay time.Duration
}

// NewButton creates a new Button struct.
func NewButton(pin machine.Pin) *Button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Button{
		Pin:           pin,
		lastInputTime: time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// Debounce overrides the default debounce delay with the provided duration value.
func (b *Button) Debounce(delay time.Duration) {
	b.debounceDelay = delay
}

// Handler sets the callback function to be call when the button is pressed.
func (b *Button) Handler(handler func(p machine.Pin)) {
	b.debounceWrapper(handler, DefaultDebounceDelay)
}

// LastInput return the time of the last button press.
func (b *Button) LastInput() time.Time {
	return b.lastInputTime
}

// Value returns true if button is currently pressed, else false.
func (b *Button) Value() bool {
	// Invert signal to match expected behavior.
	return !b.Pin.Get()
}

func (b *Button) debounceWrapper(handler func(p machine.Pin), delay time.Duration) {
	wrapped := func(p machine.Pin) {
		if time.Now().Before(b.lastInputTime.Add(delay)) {
			return
		}
		b.lastInputTime = time.Now()
		handler(p)
	}
	b.Pin.SetInterrupt(machine.PinFalling, wrapped)
}
