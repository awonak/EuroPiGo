package europi

import (
	"errors"
	"time"
)

type ApplicationStart[THardware Hardware] interface {
	Start(e THardware)
}

type ApplicationMainLoop[THardware Hardware] interface {
	MainLoop(e THardware, deltaTime time.Duration)
}

type ApplicationEnd[THardware Hardware] interface {
	End(e THardware)
}

// App sets the application handler interface with optional parameters
func App(app any, opts ...BootstrapAppOption) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if app == nil {
			return errors.New("app must not be nil")
		}

		// automatically divine the functions for the app
		start, mainLoop, end := getAppFuncs(app)

		if start == nil && mainLoop == nil && end == nil {
			return errors.New("app must provide at least one application function interface (ApplicationStart, ApplicationMainLoop, ApplicationEnd)")
		}

		o.appConfig.onAppStartFn = start
		o.appConfig.onAppMainLoopFn = mainLoop
		o.appConfig.onAppEndFn = end

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
