package europi

import (
	"errors"
	"time"
)

// MainLoopInterval sets the interval between calls to the configured main loop function
func MainLoopInterval(interval time.Duration) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if interval < 0 {
			return errors.New("interval must be greater than or equal to 0")
		}
		o.mainLoopInterval = interval
		return nil
	}
}

// EnableDisplayLogger enables (or disables) the logging of `log.Printf` (and similar) messages to
// the EuroPi's display. Enabling this will likely be undesirable except in cases where on-screen
// debugging is absoluely necessary.
func EnableDisplayLogger(enabled bool) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.enableDisplayLogger = enabled
		return nil
	}
}
