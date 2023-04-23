package rev1

import (
	"time"

	"github.com/awonak/EuroPiGo/debounce"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// digitalinput is a struct for handling reading of the digital input.
type digitalinput struct {
	dr         digitalReaderProvider
	lastChange time.Time
}

type digitalReaderProvider interface {
	Get() bool
	SetHandler(changes hal.ChangeFlags, handler func())
}

// newDigitalInput creates a new digital input struct.
func newDigitalInput(dr digitalReaderProvider) *digitalinput {
	return &digitalinput{
		dr:         dr,
		lastChange: time.Now(),
	}
}

func (d *digitalinput) Configure(config hal.DigitalInputConfig) error {
	return nil
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *digitalinput) Value() bool {
	return d.dr.Get()
}

// Handler sets the callback function to be call when the incoming signal going high event is detected.
func (d *digitalinput) Handler(handler func(value bool, deltaTime time.Duration)) {
	d.HandlerEx(hal.ChangeRising, handler)
}

// HandlerEx sets the callback function to be call when the input changes in a specified way.
func (d *digitalinput) HandlerEx(changes hal.ChangeFlags, handler func(value bool, deltaTime time.Duration)) {
	d.dr.SetHandler(changes, func() {
		now := time.Now()
		timeDelta := now.Sub(d.lastChange)
		handler(d.Value(), timeDelta)
		d.lastChange = now
	})
}

// HandlerWithDebounce sets the callback function to be call when the incoming signal going high event is detected and debounce delay time has elapsed.
func (d *digitalinput) HandlerWithDebounce(handler func(value bool, deltaTime time.Duration), delay time.Duration) {
	db := debounce.NewDebouncer(handler).Debounce(delay)
	d.Handler(func(value bool, _ time.Duration) {
		// throw away the deltaTime coming in on the handler
		// we want to use what's on the debouncer, instead
		db(value)
	})
}
