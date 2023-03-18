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
