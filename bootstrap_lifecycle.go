package europi

import "time"

func DefaultPostBootstrapInitialization(e *EuroPi) {
	e.Display.ClearBuffer()
	if err := e.Display.Display(); err != nil {
		panic(err)
	}
}

func DefaultBootstrapCompleted(e *EuroPi) {
	e.Display.ClearBuffer()
	if err := e.Display.Display(); err != nil {
		panic(err)
	}
}

// DefaultMainLoop is the default main loop used if a new one is not specified to Bootstrap()
func DefaultMainLoop(e *EuroPi, deltaTime time.Duration) {
}
