package europi

import (
	"log"
	"machine"
	"math/rand"
	"os"

	"github.com/heucuva/europi/experimental/displaylogger"
)

var (
	dispLog *displaylogger.Logger
)

func enableDisplayLogger(e *EuroPi) {
	if dispLog != nil {
		// already enabled - can happen when panicking
		return
	}

	log.SetFlags(0)
	dispLog = &displaylogger.Logger{
		Display: e.Display,
	}
	log.SetOutput(dispLog)
}

func disableDisplayLogger(e *EuroPi) {
	dispLog = nil
	log.SetOutput(os.Stdout)
}

func flushDisplayLogger(e *EuroPi) {
	if dispLog != nil {
		dispLog.Flush()
	}
}

func initRandom(e *EuroPi) {
	xl, _ := machine.GetRNG()
	xh, _ := machine.GetRNG()
	x := int64(xh)<<32 | int64(xl)
	rand.Seed(x)
}

func uninitRandom(e *EuroPi) {
}
