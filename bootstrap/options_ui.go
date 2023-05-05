package bootstrap

import (
	"errors"
	"time"

	europi "github.com/awonak/EuroPiGo"
)

// UI sets the user interface handler interface
func UI(ui UserInterface[europi.Hardware], opts ...BootstrapUIOption) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if ui == nil {
			return errors.New("ui must not be nil")
		}
		o.uiConfig.ui = ui
		o.uiConfig.options = opts
		return nil
	}
}

const (
	DefaultUIRefreshRate time.Duration = time.Millisecond * 100
)

// BootstrapOption is a single configuration parameter passed to the Bootstrap() function
type BootstrapUIOption func(o *bootstrapUIConfig) error

type bootstrapUIConfig struct {
	ui            UserInterface[europi.Hardware]
	uiRefreshRate time.Duration

	options []BootstrapUIOption
}

// UIOptions adds optional parameters for setting up the user interface
func UIOptions(option BootstrapUIOption, opts ...BootstrapUIOption) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.uiConfig.options = append(o.uiConfig.options, opts...)
		return nil
	}
}

// UIRefreshRate sets the interval of refreshes of the user interface
func UIRefreshRate(interval time.Duration) BootstrapUIOption {
	return func(o *bootstrapUIConfig) error {
		if interval <= 0 {
			return errors.New("interval must be greater than 0")
		}
		o.uiRefreshRate = interval
		return nil
	}
}
