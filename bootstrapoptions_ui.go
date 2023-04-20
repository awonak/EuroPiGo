package europi

import (
	"errors"
	"time"
)

// UI sets the user interface handler interface
func UI(ui UserInterface) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if ui == nil {
			return errors.New("ui must not be nil")
		}
		o.ui = ui
		return nil
	}
}

const (
	DefaultUIRefreshRate time.Duration = time.Millisecond * 100
)

// UIRefreshRate sets the interval of refreshes of the user interface
func UIRefreshRate(interval time.Duration) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if interval <= 0 {
			return errors.New("interval must be greater than 0")
		}
		o.uiRefreshRate = interval
		return nil
	}
}
