package europi

import (
	"time"
)

// BootstrapOption is a single configuration parameter passed to the Bootstrap() function
type BootstrapOption func(o *bootstrapConfig) error

type bootstrapConfig struct {
	appMainLoopInterval    time.Duration
	panicHandler           func(e *EuroPi, reason any)
	enableDisplayLogger    bool
	initRandom             bool
	europi                 *EuroPi
	enableNonPicoWebSocket bool

	// user interface
	ui            UserInterface
	uiRefreshRate time.Duration

	// lifecycle callbacks
	onPostBootstrapConstructionFn PostBootstrapConstructionFunc
	onPreInitializeComponentsFn   PreInitializeComponentsFunc
	onPostInitializeComponentsFn  PostInitializeComponentsFunc
	onBootstrapCompletedFn        BootstrapCompletedFunc
	onAppStartFn                  AppStartFunc
	onAppMainLoopFn               AppMainLoopFunc
	onAppEndFn                    AppEndFunc
	onBeginDestroyFn              BeginDestroyFunc
	onFinishDestroyFn             FinishDestroyFunc
}
