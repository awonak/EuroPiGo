package bootstrap

import (
	"context"
	"errors"
	"sync"
	"time"

	europi "github.com/awonak/EuroPiGo"
)

// Bootstrap will set up a global runtime environment (see europi.Pi)
func Bootstrap(pi europi.Hardware, options ...BootstrapOption) error {
	e := pi
	if e == nil {
		return errors.New("europi must be provided")
	}

	config := bootstrapConfig{
		panicHandler:           DefaultPanicHandler,
		enableDisplayLogger:    DefaultEnableDisplayLogger,
		initRandom:             DefaultInitRandom,
		enableNonPicoWebSocket: defaultWebSimEnabled,
		europi:                 e,

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

	var (
		onceBootstrapDestroy sync.Once
		nonPicoWSApi         NonPicoWSActivation
	)
	panicHandler := config.panicHandler
	lastDestroyFunc := config.onBeginDestroyFn
	ctx := e.Context()
	runBootstrapDestroy := func() {
		reason := recover()
		_ = e.Shutdown(reason)
		if reason != nil && panicHandler != nil {
			config.onBeginDestroyFn = func(e europi.Hardware, reason any) {
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

	nonPicoWSApi = bootstrapInitializeComponents(ctx, &config, e)

	if config.onBootstrapCompletedFn != nil {
		config.onBootstrapCompletedFn(e)
	}

	bootstrapRunLoop(&config, e)

	return nil
}

func Shutdown(e europi.Hardware, reason any) error {
	return e.Shutdown(reason)
}

func bootstrapInitializeComponents(ctx context.Context, config *bootstrapConfig, e europi.Hardware) NonPicoWSActivation {
	if config.onPreInitializeComponentsFn != nil {
		config.onPreInitializeComponentsFn(e)
	}

	if config.enableDisplayLogger {
		enableDisplayLogger(e)
	}

	var nonPicoWSApi NonPicoWSActivation
	if config.enableNonPicoWebSocket {
		nonPicoWSApi = ActivateNonPicoWS(ctx, e)
	}

	if config.initRandom {
		initRandom(e)
	}

	// ui initializaiton is always last
	if config.uiConfig.ui != nil {
		enableUI(ctx, e, config.uiConfig)
	}

	if config.onPostInitializeComponentsFn != nil {
		config.onPostInitializeComponentsFn(e)
	}

	return nonPicoWSApi
}

func bootstrapRunLoop(config *bootstrapConfig, e europi.Hardware) {
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

func bootstrapRunLoopWithDelay(config *bootstrapConfig, e europi.Hardware) {
	if config.appConfig.onAppMainLoopFn == nil {
		panic(errors.New("no main loop specified"))
	}

	ticker := time.NewTicker(config.appConfig.mainLoopInterval)
	defer ticker.Stop()

	lastTick := time.Now()
	for {
		select {
		case reason := <-e.Context().Done():
			panic(reason)

		case now := <-ticker.C:
			config.appConfig.onAppMainLoopFn(e, now.Sub(lastTick))
			lastTick = now
		}
	}
}

func bootstrapRunLoopNoDelay(config *bootstrapConfig, e europi.Hardware) {
	if config.appConfig.onAppMainLoopFn == nil {
		panic(errors.New("no main loop specified"))
	}

	lastTick := time.Now()
	for {
		select {
		case reason := <-e.Context().Done():
			panic(reason)

		default:
			now := time.Now()
			config.appConfig.onAppMainLoopFn(e, now.Sub(lastTick))
			lastTick = now
		}
	}
}

func bootstrapDestroy(config *bootstrapConfig, e europi.Hardware, nonPicoWSApi NonPicoWSActivation, reason any) {
	if config.onBeginDestroyFn != nil {
		config.onBeginDestroyFn(e, reason)
	}

	disableUI(e)

	if config.enableNonPicoWebSocket && deactivateNonPicoWebSocket != nil {
		deactivateNonPicoWebSocket(e, nonPicoWSApi)
	}

	disableDisplayLogger(e)

	uninitRandom(e)

	if display := europi.Display(e); display != nil {
		// show the last buffer
		_ = display.Display()
	}

	if config.onFinishDestroyFn != nil {
		config.onFinishDestroyFn(e)
	}
}
