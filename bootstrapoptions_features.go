package europi

import (
	"errors"
	"time"
)

const (
	DefaultAppMainLoopInterval time.Duration = time.Millisecond * 100
)

// AppMainLoopInterval sets the interval between calls to the configured app main loop function
func AppMainLoopInterval(interval time.Duration) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if interval < 0 {
			return errors.New("interval must be greater than or equal to 0")
		}
		o.appMainLoopInterval = interval
		return nil
	}
}

const (
	DefaultEnableDisplayLogger bool = false
)

// EnableDisplayLogger enables (or disables) the logging of `log.Printf` (and similar) messages to
// the EuroPi's display. Enabling this will likely be undesirable except in cases where on-screen
// debugging is absoluely necessary.
func EnableDisplayLogger(enabled bool) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.enableDisplayLogger = enabled
		return nil
	}
}

const (
	DefaultInitRandom bool = true
)

// InitRandom enables (or disables) the initialization of the Go standard library's `rand` package
// Seed value. Disabling this will likely be undesirable except in cases where deterministic 'random'
// number generation is required, as the standard library `rand` package defaults to a seed of 1
// instead of some pseudo-random number, like current time or thermal values.
// To generate a pseudo-random number for the random seed, the `machine.GetRNG` function is used.
func InitRandom(enabled bool) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.initRandom = enabled
		return nil
	}
}

// UsingEuroPi sets a specific EuroPi object instance for all operations in the bootstrap
func UsingEuroPi(e *EuroPi) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if e == nil {
			return errors.New("europi instance must not be nil")
		}

		o.europi = e
		return nil
	}
}

// AttachNonPicoWS (if enabled and on non-Pico builds with build flags of `-tags=revision1` set)
// starts up a websocket interface and system debugger on port 8080
func AttachNonPicoWS(enabled bool) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.enableNonPicoWebSocket = enabled
		return nil
	}
}
