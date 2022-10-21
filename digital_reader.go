package europi

import (
	"machine"
	"time"
)

const defaultDebounceDelay = time.Duration(50 * time.Millisecond)

// DigitalReader is an interface for common digital inputs methods.
type DigitalReader interface {
	Handler(func(machine.Pin))
	HandlerWithDebounce(func(machine.Pin), time.Duration)
	LastInput() time.Time
	Value() bool
}

type digitalReader struct {
	pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(machine.Pin)
}

func newDigitalReader(pin machine.Pin) *digitalReader {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &digitalReader{
		pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: defaultDebounceDelay,
	}
}

// LastInput return the time of the last high input (triggered at 0.8v).
func (d *digitalReader) LastInput() time.Time {
	return d.lastInput
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *digitalReader) Value() bool {
	// Invert signal to match expected behavior.
	return !d.pin.Get()
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *digitalReader) Handler(callback func(machine.Pin)) {
	d.HandlerWithDebounce(callback, 0)
}

// Handler sets the callback function to be call when a rising edge is detected and debounce delay time has elapsed.
func (d *digitalReader) HandlerWithDebounce(callback func(machine.Pin), delay time.Duration) {
	d.callback = callback
	d.debounceDelay = delay
	d.pin.SetInterrupt(machine.PinFalling, d.debounceWrapper)
}

func (d *digitalReader) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(d.lastInput.Add(d.debounceDelay)) {
		return
	}
	d.callback(p)
	d.lastInput = t
}

type digitalInput struct {
	DigitalReader
}

func newDigitalInput(pin machine.Pin) *digitalInput {
	return &digitalInput{
		newDigitalReader(pin),
	}
}

type button struct {
	DigitalReader
}

func newButton(pin machine.Pin) *button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &button{
		newDigitalReader(pin),
	}
}
