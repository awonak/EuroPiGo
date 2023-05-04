package europi

import "time"

func DefaultPostBootstrapInitialization(e Hardware) {
	display := Display(e)
	if display == nil {
		// no display, can't continue
		return
	}

	display.ClearBuffer()
	if err := display.Display(); err != nil {
		panic(err)
	}
}

func DefaultBootstrapCompleted(e Hardware) {
	display := Display(e)
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
func DefaultMainLoop(e Hardware, deltaTime time.Duration) {
}
