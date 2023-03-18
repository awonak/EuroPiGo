package europi

import (
	"errors"
	"time"
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
	Callback: StartLoop
		|
		V
	Callback(on tick): MainLoop
		|
		V
	Callback: EndLoop
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
	PostBootstrapConstructionFunc func(e *EuroPi)
	PreInitializeComponentsFunc   func(e *EuroPi)
	PostInitializeComponentsFunc  func(e *EuroPi)
	BootstrapCompletedFunc        func(e *EuroPi)
	StartLoopFunc                 func(e *EuroPi)
	MainLoopFunc                  func(e *EuroPi, deltaTime time.Duration)
	EndLoopFunc                   func(e *EuroPi)
	BeginDestroyFunc              func(e *EuroPi)
	FinishDestroyFunc             func(e *EuroPi)
)

// PostBootstrapConstruction runs immediately after primary EuroPi bootstrap has finished,
// but before components have been initialized
func PostBootstrapConstruction(fn PostBootstrapConstructionFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onPostBootstrapConstructionFn = fn
		return nil
	}
}

func PreInitializeComponents(fn PreInitializeComponentsFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onPreInitializeComponentsFn = fn
		return nil
	}
}

func PostInitializeComponents(fn PostInitializeComponentsFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onPostInitializeComponentsFn = fn
		return nil
	}
}

func BootstrapCompleted(fn BootstrapCompletedFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onBootstrapCompletedFn = fn
		return nil
	}
}

func StartLoop(fn StartLoopFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onStartLoopFn = fn
		return nil
	}
}

// MainLoop sets the main loop function to be called on interval.
// nil is not allowed - if you want to set the default, either do not specify a MainLoop() option
// or specify europi.DefaultMainLoop
func MainLoop(fn MainLoopFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		if fn == nil {
			return errors.New("a valid main loop function must be specified")
		}
		o.onMainLoopFn = fn
		return nil
	}
}

func EndLoop(fn EndLoopFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onEndLoopFn = fn
		return nil
	}
}

func BeginDestroy(fn BeginDestroyFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onBeginDestroyFn = fn
		return nil
	}
}

func FinishDestroy(fn FinishDestroyFunc) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.onFinishDestroyFn = fn
		return nil
	}
}
