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

// Handler sets the callback function to be call when a rising edge is detected.
func (d *DigitalInput) Handler(handler func(p machine.Pin)) {
	d.Pin.SetInterrupt(machine.PinRising, handler)
}

// HandlerWithDebounce sets the callback function to be call when a rising edge
// is detected and it has been `delay` duration since last rising edge.
func (d *DigitalInput) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	d.Pin.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		t := time.Now()
		if t.Before(d.lastInputTime.Add(delay)) {
			return
		}
		handler(p)
		d.lastInputTime = t
	})
}

// Button is a struct for handling push button behavior.
type Button struct {
	Pin           machine.Pin
	lastInputTime time.Time
	debounceDelay time.Duration
	callback      func(p machine.Pin)
}

// NewButton creates a new Button struct.
func NewButton(pin machine.Pin) *Button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Button{
		Pin:           pin,
		lastInputTime: time.Now(),
	}
}

// Handler sets the callback function to be call when the button is pressed.
func (b *Button) Handler(handler func(p machine.Pin)) {
	b.Pin.SetInterrupt(machine.PinFalling, handler)
}

// Handler sets the callback function to be call when the button is pressed.
func (b *Button) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	b.callback = handler
	b.debounceDelay = delay
	b.Pin.SetInterrupt(machine.PinFalling, b.debounceWrapper)
}

func (b *Button) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(b.lastInputTime.Add(b.debounceDelay)) {
		return
	}
	b.callback(p)
	b.lastInputTime = t
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
