package europi

import (
	"context"
	"log"
	"os"

	"github.com/awonak/EuroPiGo/experimental/displaylogger"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

var (
	dispLog displaylogger.Logger
)

func enableDisplayLogger(e Hardware) {
	if dispLog != nil {
		// already enabled - can happen when panicking
		return
	}

	display := Display(e)
	if display == nil {
		// no display, can't continue
		return
	}

	log.SetFlags(0)
	dispLog = displaylogger.NewLogger(display)
	log.SetOutput(dispLog)
}

func disableDisplayLogger(e Hardware) {
	flushDisplayLogger(e)
	dispLog = nil
	log.SetOutput(os.Stdout)
}

func flushDisplayLogger(e Hardware) {
	if dispLog != nil {
		dispLog.Flush()
	}
}

func initRandom(e Hardware) {
	if rnd := e.Random(); rnd != nil {
		_ = rnd.Configure(hal.RandomGeneratorConfig{})
	}
}

func uninitRandom(e Hardware) {
}

// used for non-pico testing of bootstrapped europi apps
var (
	activateNonPicoWebSocket   func(ctx context.Context, e Hardware) NonPicoWSActivation
	deactivateNonPicoWebSocket func(e Hardware, api NonPicoWSActivation)
)

type NonPicoWSActivation interface {
	Shutdown() error
}

func ActivateNonPicoWS(ctx context.Context, e Hardware) NonPicoWSActivation {
	if activateNonPicoWebSocket == nil {
		return nil
	}
	return activateNonPicoWebSocket(ctx, e)
}
