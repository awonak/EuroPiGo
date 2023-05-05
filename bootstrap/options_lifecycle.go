package bootstrap

import (
	"errors"
	"time"

	europi "github.com/awonak/EuroPiGo"
)

/* Order of lifecycle calls:
BootStrap
	|
	V
Callback: PostBootstrapConstruction
	|
	V
Bootstrap: postBootstrapConstruction
		|
		V
	Callback: PreInitializeComponents
		|
		V
	Bootstrap: initializeComponents
		|
		V
	Callback: PostInitializeComponents
		|
		V
Callback: BootstrapCompleted
	|
	V
Bootstrap: runLoop
		|
		V
	Callback: AppStart
		|
		V
	Callback(on tick): AppMainLoop
		|
		V
	Callback: AppEnd
		|
		V
Bootstrap: destroyBootstrap
	    |
		V
	Callback: BeginDestroy
		|
		V
	Callback: FinishDestroy
*/

type (
	PostBootstrapConstructionFunc func(e europi.Hardware)
	PreInitializeComponentsFunc   func(e europi.Hardware)
	PostInitializeComponentsFunc  func(e europi.Hardware)
	BootstrapCompletedFunc        func(e europi.Hardware)
	AppStartFunc                  func(e europi.Hardware)
	AppMainLoopFunc               func(e europi.Hardware, deltaTime time.Duration)
	AppEndFunc                    func(e europi.Hardware)
	BeginDestroyFunc              func(e europi.Hardware, reason any)
	FinishDestroyFunc             func(e europi.Hardware)
)

// PostBootstrapConstruction sets the function that runs immediately after primary EuroPi bootstrap
// has finished, but before components have been initialized. Nearly none of the functionality of
// the bootstrap is ready or configured at this point.
func PostBootstrapConstruction(fn PostBootstrapConstructionFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onPostBootstrapConstructionFn = fn
		return nil
	}
}

// PreInitializeComponents sets the function that recevies notification of when components of the
// bootstrap are about to start their initialization phase and the bootstrap is getting ready.
// Most operational functionality of the bootstrap is definitely not configured at this point.
func PreInitializeComponents(fn PreInitializeComponentsFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onPreInitializeComponentsFn = fn
		return nil
	}
}

// PostInitializeComponents sets the function that recevies notification of when components of the
// bootstrap have completed their initialization phase and the bootstrap is nearly ready for full
// operation. Some operational functionality of the bootstrap might not be configured at this point.
func PostInitializeComponents(fn PostInitializeComponentsFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onPostInitializeComponentsFn = fn
		return nil
	}
}

// BootstrapCompleted sets the function that receives notification of critical bootstrap
// operations being complete - this is the first point where functions within the bootstrap
// may be used without fear of there being an incomplete operating state.
func BootstrapCompleted(fn BootstrapCompletedFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onBootstrapCompletedFn = fn
		return nil
	}
}

// TODO: consider secondary bootloader support functionality here once internal flash support
// becomes a reality.

// AppStart sets the application function to be called before the main operating loop
// processing begins. At this point, the bootstrap configuration has completed and
// all bootstrap functionality may be used without fear of there being an incomplete
// operating state.
func AppStart(fn AppStartFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.appConfig.onAppStartFn = fn
		return nil
	}
}

// AppMainLoop sets the application main loop function to be called on interval.
// nil is not allowed - if you want to set the default, either do not specify a AppMainLoop() option
// or specify europi.DefaultMainLoop
func AppMainLoop(fn AppMainLoopFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if fn == nil {
			return errors.New("a valid main loop function must be specified")
		}
		o.appConfig.onAppMainLoopFn = fn
		return nil
	}
}

// AppEnd sets the application function that's called right before the bootstrap
// destruction processing is performed.
func AppEnd(fn AppEndFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.appConfig.onAppEndFn = fn
		return nil
	}
}

// BeginDestroy sets the function that receives the notification of shutdown of the bootstrap and
// is also the first stop within the `panic()` handler functionality. If the `reason` parameter
// is non-nil, then a critical failure has been detected and the bootstrap is in the last stages of
// complete destruction. If it is nil, then it can be assumed that proper functionality of the
// bootstrap is still available, but heading towards the last steps of unavailability once the
// function exits.
func BeginDestroy(fn BeginDestroyFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onBeginDestroyFn = fn
		return nil
	}
}

// FinishDestroy sets the function that receives the final notification of shutdown of the bootstrap.
// The entire bootstrap is disabled, all timers, queues, and components are considered deactivated.
func FinishDestroy(fn FinishDestroyFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onFinishDestroyFn = fn
		return nil
	}
}
