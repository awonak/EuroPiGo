package europi

import (
	"errors"
	"time"
)

type ApplicationStart interface {
	Start(e *EuroPi)
}

type ApplicationMainLoop interface {
	MainLoop(e *EuroPi, deltaTime time.Duration)
}

type ApplicationEnd interface {
	End(e *EuroPi)
}

// App sets the application handler interface with optional parameters
func App(app any, opts ...BootstrapAppOption) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if app == nil {
			return errors.New("app must not be nil")
		}
		start, _ := app.(ApplicationStart)
		mainLoop, _ := app.(ApplicationMainLoop)
		end, _ := app.(ApplicationEnd)

		if start == nil && mainLoop == nil && end == nil {
			return errors.New("app must provide at least one application function interface (ApplicationStart, ApplicationMainLoop, ApplicationEnd)")
		}

		if start != nil {
			o.appConfig.onAppStartFn = start.Start
		}
		if mainLoop != nil {
			o.appConfig.onAppMainLoopFn = mainLoop.MainLoop
		}
		if end != nil {
			o.appConfig.onAppEndFn = end.End
		}

		o.appConfig.options = opts
		return nil
	}
}

// AppOptions adds optional parameters for setting up the application interface
func AppOptions(option BootstrapAppOption, opts ...BootstrapAppOption) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.appConfig.options = append(o.appConfig.options, opts...)
		return nil
	}
}

// BootstrapAppOption is a single configuration parameter passed to the App() or AppOption() functions
type BootstrapAppOption func(o *bootstrapAppConfig) error

type bootstrapAppConfig struct {
	mainLoopInterval time.Duration
	onAppStartFn     AppStartFunc
	onAppMainLoopFn  AppMainLoopFunc
	onAppEndFn       AppEndFunc

	options []BootstrapAppOption
}

const (
	DefaultAppMainLoopInterval time.Duration = time.Millisecond * 100
)

// AppMainLoopInterval sets the interval between calls to the configured app main loop function
func AppMainLoopInterval(interval time.Duration) BootstrapAppOption {
	return func(o *bootstrapAppConfig) error {
		if interval < 0 {
			return errors.New("interval must be greater than or equal to 0")
		}
		o.mainLoopInterval = interval
		return nil
	}
}
