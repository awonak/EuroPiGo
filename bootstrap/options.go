package bootstrap

import europi "github.com/awonak/EuroPiGo"

// BootstrapOption is a single configuration parameter passed to the Bootstrap() function
type BootstrapOption func(o *bootstrapConfig) error

type bootstrapConfig struct {
	panicHandler           func(e europi.Hardware, reason any)
	enableDisplayLogger    bool
	initRandom             bool
	europi                 europi.Hardware
	enableNonPicoWebSocket bool

	// application
	appConfig bootstrapAppConfig

	// user interface
	uiConfig bootstrapUIConfig

	// lifecycle callbacks
	onPostBootstrapConstructionFn PostBootstrapConstructionFunc
	onPreInitializeComponentsFn   PreInitializeComponentsFunc
	onPostInitializeComponentsFn  PostInitializeComponentsFunc
	onBootstrapCompletedFn        BootstrapCompletedFunc
	onBeginDestroyFn              BeginDestroyFunc
	onFinishDestroyFn             FinishDestroyFunc
}
