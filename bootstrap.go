package europi

import (
	"errors"
	"sync"
	"time"
)

var (
	// Pi is a global EuroPi instance constructed by calling the Bootstrap() function
	Pi *EuroPi

	piWantDestroyChan chan any
)

// Bootstrap will set up a global runtime environment (see europi.Pi)
func Bootstrap(options ...BootstrapOption) error {
	config := bootstrapConfig{
		panicHandler:           DefaultPanicHandler,
		enableDisplayLogger:    DefaultEnableDisplayLogger,
		initRandom:             DefaultInitRandom,
		enableNonPicoWebSocket: false,
		europi:                 nil,

		appConfig: bootstrapAppConfig{
			mainLoopInterval: DefaultAppMainLoopInterval,
			onAppStartFn:     nil,
			onAppMainLoopFn:  DefaultMainLoop,
			onAppEndFn:       nil,
		},

		uiConfig: bootstrapUIConfig{
			ui:            nil,
			uiRefreshRate: DefaultUIRefreshRate,
		},

		onPostBootstrapConstructionFn: DefaultPostBootstrapInitialization,
		onPreInitializeComponentsFn:   nil,
		onPostInitializeComponentsFn:  nil,
		onBootstrapCompletedFn:        DefaultBootstrapCompleted,
		onBeginDestroyFn:              nil,
		onFinishDestroyFn:             nil,
	}

	// process bootstrap options
	for _, opt := range options {
		if err := opt(&config); err != nil {
			return err
		}
	}

	// process app options
	for _, opt := range config.appConfig.options {
		if err := opt(&config.appConfig); err != nil {
			return err
		}
	}

	// process ui options
	for _, opt := range config.uiConfig.options {
		if err := opt(&config.uiConfig); err != nil {
			return err
		}
	}

	if config.europi == nil {
		config.europi = New()
	}
	e := config.europi

	if e == nil {
		return errors.New("no europi available")
	}

	Pi = e
	piWantDestroyChan = make(chan any, 1)

	var (
		onceBootstrapDestroy sync.Once
		nonPicoWSApi         nonPicoWSActivation
	)
	panicHandler := config.panicHandler
	lastDestroyFunc := config.onBeginDestroyFn
	runBootstrapDestroy := func() {
		reason := recover()
		if reason != nil && panicHandler != nil {
			config.onBeginDestroyFn = func(e *EuroPi, reason any) {
				if lastDestroyFunc != nil {
					lastDestroyFunc(e, reason)
				}
				panicHandler(e, reason)
			}
		}
		onceBootstrapDestroy.Do(func() {
			bootstrapDestroy(&config, e, nonPicoWSApi, reason)
		})
	}
	defer runBootstrapDestroy()

	if config.onPostBootstrapConstructionFn != nil {
		config.onPostBootstrapConstructionFn(e)
	}

	nonPicoWSApi = bootstrapInitializeComponents(&config, e)

	if config.onBootstrapCompletedFn != nil {
		config.onBootstrapCompletedFn(e)
	}

	bootstrapRunLoop(&config, e)

	return nil
}

func Shutdown(reason any) error {
	if piWantDestroyChan == nil {
		return errors.New("cannot shutdown: no available bootstrap")
	}

	piWantDestroyChan <- reason
	return nil
}

func bootstrapInitializeComponents(config *bootstrapConfig, e *EuroPi) nonPicoWSActivation {
	if config.onPreInitializeComponentsFn != nil {
		config.onPreInitializeComponentsFn(e)
	}

	if config.enableDisplayLogger {
		enableDisplayLogger(e)
	}

	var nonPicoWSApi nonPicoWSActivation
	if config.enableNonPicoWebSocket && activateNonPicoWebSocket != nil {
		nonPicoWSApi = activateNonPicoWebSocket(e)
	}

	if config.initRandom {
		initRandom(e)
	}

	// ui initializaiton is always last
	if config.uiConfig.ui != nil {
		enableUI(e, config.uiConfig)
	}

	if config.onPostInitializeComponentsFn != nil {
		config.onPostInitializeComponentsFn(e)
	}

	return nonPicoWSApi
}

func bootstrapRunLoop(config *bootstrapConfig, e *EuroPi) {
	if config.appConfig.onAppStartFn != nil {
		config.appConfig.onAppStartFn(e)
	}

	startUI(e)

	ForceRepaintUI(e)

	if config.appConfig.mainLoopInterval > 0 {
		bootstrapRunLoopWithDelay(config, e)
	} else {
		bootstrapRunLoopNoDelay(config, e)
	}

	if config.appConfig.onAppEndFn != nil {
		config.appConfig.onAppEndFn(e)
	}
}

func bootstrapRunLoopWithDelay(config *bootstrapConfig, e *EuroPi) {
	if config.appConfig.onAppMainLoopFn == nil {
		panic(errors.New("no main loop specified"))
	}

	ticker := time.NewTicker(config.appConfig.mainLoopInterval)
	defer ticker.Stop()

	lastTick := time.Now()
	for {
		select {
		case reason := <-piWantDestroyChan:
			panic(reason)

		case now := <-ticker.C:
			config.appConfig.onAppMainLoopFn(e, now.Sub(lastTick))
			lastTick = now
		}
	}
}

func bootstrapRunLoopNoDelay(config *bootstrapConfig, e *EuroPi) {
	if config.appConfig.onAppMainLoopFn == nil {
		panic(errors.New("no main loop specified"))
	}

	lastTick := time.Now()
	for {
		select {
		case reason := <-piWantDestroyChan:
			panic(reason)

		default:
			now := time.Now()
			config.appConfig.onAppMainLoopFn(e, now.Sub(lastTick))
			lastTick = now
		}
	}
}

func bootstrapDestroy(config *bootstrapConfig, e *EuroPi, nonPicoWSApi nonPicoWSActivation, reason any) {
	if config.onBeginDestroyFn != nil {
		config.onBeginDestroyFn(e, reason)
	}

	disableUI(e)

	if config.enableNonPicoWebSocket && deactivateNonPicoWebSocket != nil {
		deactivateNonPicoWebSocket(e, nonPicoWSApi)
	}

	disableDisplayLogger(e)

	uninitRandom(e)

	if e != nil && e.Display != nil {
		// show the last buffer
		_ = e.Display.Display()
	}

	close(piWantDestroyChan)
	Pi = nil

	if config.onFinishDestroyFn != nil {
		config.onFinishDestroyFn(e)
	}
}
