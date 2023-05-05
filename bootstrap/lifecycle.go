package bootstrap

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
)

func DefaultPostBootstrapInitialization(e europi.Hardware) {
	display := europi.Display(e)
	if display == nil {
		// no display, can't continue
		return
	}

	display.ClearBuffer()
	if err := display.Display(); err != nil {
		panic(err)
	}
}

func DefaultBootstrapCompleted(e europi.Hardware) {
	display := europi.Display(e)
	if display == nil {
		// no display, can't continue
		return
	}

	display.ClearBuffer()
	if err := display.Display(); err != nil {
		panic(err)
	}
}

// DefaultMainLoop is the default main loop used if a new one is not specified to Bootstrap()
func DefaultMainLoop(e europi.Hardware, deltaTime time.Duration) {
}
