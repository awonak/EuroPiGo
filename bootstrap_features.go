package europi

import (
	"log"
	"os"

	"github.com/heucuva/europi/experimental/displaylogger"
	"github.com/heucuva/europi/internal/hardware/hal"
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
		e.RND.Configure(hal.RandomGeneratorConfig{})
	}
}

func uninitRandom(e *EuroPi) {
}
