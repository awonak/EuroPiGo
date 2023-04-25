package europi

import (
	"log"
	"os"

	"github.com/awonak/EuroPiGo/experimental/displaylogger"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

var (
	dispLog displaylogger.Logger
)

func enableDisplayLogger(e *EuroPi) {
	if dispLog != nil {
		// already enabled - can happen when panicking
		return
	}

	log.SetFlags(0)
	dispLog = displaylogger.NewLogger(e.Display)
	log.SetOutput(dispLog)
}

func disableDisplayLogger(e *EuroPi) {
	flushDisplayLogger(e)
	dispLog = nil
	log.SetOutput(os.Stdout)
}

func flushDisplayLogger(e *EuroPi) {
	if dispLog != nil {
		dispLog.Flush()
	}
}

func initRandom(e *EuroPi) {
	if e.RND != nil {
		_ = e.RND.Configure(hal.RandomGeneratorConfig{})
	}
}

func uninitRandom(e *EuroPi) {
}

// used for non-pico testing of bootstrapped europi apps
var (
	activateNonPicoWebSocket   func(e *EuroPi) nonPicoWSActivation
	deactivateNonPicoWebSocket func(e *EuroPi, api nonPicoWSActivation)
)

type nonPicoWSActivation interface {
	Shutdown() error
}
