package europi

import (
	"machine"
	"time"
)

const defaultDebounceDelay = time.Duration(50 * time.Millisecond)

type callback func(machine.Pin)

// DigitalReader is an interface for common digital inputs methods.
type DigitalReader interface {
	Handler(callback)
	HandlerWithDebounce(callback, time.Duration)
	LastInput() time.Time
	Value() bool
}

type digitalReader struct {
	machine.Pin

	debounceDelay time.Duration
	lastInput     time.Time
	callback      callback
}

func newDigitalReader(pin machine.Pin) *digitalReader {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &digitalReader{
		Pin:           pin,
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
	return !d.Pin.Get()
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *digitalReader) Handler(c callback) {
	d.HandlerWithDebounce(c, 0)
}

// Handler sets the callback function to be call when a rising edge is detected and debounce delay time has elapsed.
func (d *digitalReader) HandlerWithDebounce(c callback, delay time.Duration) {
	d.callback = c
	d.debounceDelay = delay
	d.Pin.SetInterrupt(machine.PinFalling, d.debounceWrapper)
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
