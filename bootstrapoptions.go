package europi

import (
	"time"
)

// BootstrapOption is a single configuration parameter passed to the Bootstrap() function
type BootstrapOption func(o *bootstrapConfig) error

type bootstrapConfig struct {
	mainLoopInterval time.Duration

	// lifecycle callbacks
	onPostBootstrapConstructionFn PostBootstrapConstructionFunc
	onPreInitializeComponentsFn   PreInitializeComponentsFunc
	onPostInitializeComponentsFn  PostInitializeComponentsFunc
	onBootstrapCompletedFn        BootstrapCompletedFunc
	onStartLoopFn                 StartLoopFunc
	onMainLoopFn                  MainLoopFunc
	onEndLoopFn                   EndLoopFunc
	onBeginDestroyFn              BeginDestroyFunc
	onFinishDestroyFn             FinishDestroyFunc
}
