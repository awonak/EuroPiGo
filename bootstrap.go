package europi

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/heucuva/europi/output"
)

var (
	// Pi is a global EuroPi instance constructed by calling the Bootstrap() function
	Pi *EuroPi

	piWantDestroyChan chan struct{}
)

// Bootstrap will set up a global runtime environment (see europi.Pi)
func Bootstrap(options ...BootstrapOption) error {
	config := bootstrapConfig{
		mainLoopInterval: DefaultMainLoopInterval,

		onPostBootstrapConstructionFn: DefaultPostBootstrapInitialization,
		onPreInitializeComponentsFn:   nil,
		onPostInitializeComponentsFn:  nil,
		onBootstrapCompletedFn:        DefaultBootstrapCompleted,
		onStartLoopFn:                 nil,
		onMainLoopFn:                  DefaultMainLoop,
		onEndLoopFn:                   nil,
		onBeginDestroyFn:              nil,
		onFinishDestroyFn:             nil,
	}

	for _, opt := range options {
		if err := opt(&config); err != nil {
			return err
		}
	}

	e := New()

	Pi = e
	piWantDestroyChan = make(chan struct{}, 1)

	defer func() {
		if err := recover(); err != nil {
			fnt := output.DefaultFont
			e.Display.SetFont(fnt)
			e.Display.WriteLine(fmt.Sprint(err), 0, int16(fnt.YAdvance))
			_ = e.Display.Display()
		}
	}()

	var onceBootstrapDestroy sync.Once
	runBootstrapDestroy := func() {
		onceBootstrapDestroy.Do(func() {
			bootstrapDestroy(&config, e)
		})
	}
	defer runBootstrapDestroy()

	if config.onPostBootstrapConstructionFn != nil {
		config.onPostBootstrapConstructionFn(e)
	}

	bootstrapInitializeComponents(&config, e)

	if config.onBootstrapCompletedFn != nil {
		config.onBootstrapCompletedFn(e)
	}

	bootstrapRunLoop(&config, e)

	return nil
}

func Shutdown() error {
	if piWantDestroyChan == nil {
		return errors.New("cannot shutdown: no available bootstrap")
	}

	piWantDestroyChan <- struct{}{}
	return nil
}

func bootstrapInitializeComponents(config *bootstrapConfig, e *EuroPi) {
	if config.onPreInitializeComponentsFn != nil {
		config.onPreInitializeComponentsFn(e)
	}

	// TODO: initialize components

	if config.onPostInitializeComponentsFn != nil {
		config.onPostInitializeComponentsFn(e)
	}
}

func bootstrapRunLoop(config *bootstrapConfig, e *EuroPi) {
	if config.onStartLoopFn != nil {
		config.onStartLoopFn(e)
	}

	if config.mainLoopInterval > 0 {
		bootstrapRunLoopWithDelay(config, e)
	} else {
		bootstrapRunLoopNoDelay(config, e)
	}

	if config.onEndLoopFn != nil {
		config.onEndLoopFn(e)
	}
}

func bootstrapRunLoopWithDelay(config *bootstrapConfig, e *EuroPi) {
	ticker := time.NewTicker(config.mainLoopInterval)
	defer ticker.Stop()

	lastTick := time.Now()
mainLoop:
	for {
		select {
		case <-piWantDestroyChan:
			break mainLoop

		case now := <-ticker.C:
			config.onMainLoopFn(e, now.Sub(lastTick))
			lastTick = now
		}
	}
}

func bootstrapRunLoopNoDelay(config *bootstrapConfig, e *EuroPi) {
	lastTick := time.Now()
mainLoop:
	for {
		select {
		case <-piWantDestroyChan:
			break mainLoop

		default:
			now := time.Now()
			config.onMainLoopFn(e, now.Sub(lastTick))
			lastTick = now
		}
	}
}

func bootstrapDestroy(config *bootstrapConfig, e *EuroPi) {
	if config.onBeginDestroyFn != nil {
		config.onBeginDestroyFn(e)
	}

	close(piWantDestroyChan)
	Pi = nil

	if config.onFinishDestroyFn != nil {
		config.onFinishDestroyFn(e)
	}
}
