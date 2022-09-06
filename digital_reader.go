package europi

import (
	"fmt"
	"machine"
	"time"
)

var DefaultDebounceDelay time.Duration

func init() {
	var err error
	DefaultDebounceDelay, err = time.ParseDuration("500ms")
	if err != nil {
		fmt.Print(fmt.Errorf("defaultDebounce fail: %e", err))
	}
}

type DigitalReader interface {
	Handler(func(machine.Pin))
	LastInput() time.Time
	Value() bool
}

type DigitalInput struct {
	Pin           machine.Pin
	lastInputTime time.Time
}

func NewDI(pin machine.Pin) *DigitalInput {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &DigitalInput{
		Pin:           pin,
		lastInputTime: time.Now(),
	}
}

func (d *DigitalInput) LastInput() time.Time {
	return d.lastInputTime
}

func (d *DigitalInput) Value() bool {
	// Invert signal to match expected behavior.
	return !d.Pin.Get()
}

func (d *DigitalInput) Handler(handler func(p machine.Pin)) {
	d.debounceWrapper(handler, DefaultDebounceDelay)
}
func (d *DigitalInput) debounceWrapper(handler func(p machine.Pin), delay time.Duration) {
	wrapped := func(p machine.Pin) {
		if time.Now().Before(d.lastInputTime.Add(delay)) {
			return
		}
		d.lastInputTime = time.Now()
		handler(p)
	}
	d.Pin.SetInterrupt(machine.PinFalling, wrapped)
}

type Button struct {
	Pin           machine.Pin
	lastInputTime time.Time
}

func NewButton(pin machine.Pin) *Button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Button{
		Pin:           pin,
		lastInputTime: time.Now(),
	}
}

func (b *Button) Handler(handler func(p machine.Pin)) {
	b.debounceWrapper(handler, DefaultDebounceDelay)
}

func (b *Button) LastInput() time.Time {
	return b.lastInputTime
}

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
