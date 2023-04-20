package europi

import (
	"fmt"
	"log"

	"github.com/heucuva/europi/output"
)

// DefaultPanicHandler is the default handler for panics
// This will be set by the build flag `onscreenpanic` to `handlePanicOnScreenLog`
// Not setting the build flag will set it to `handlePanicDisplayCrash`
var DefaultPanicHandler func(e *EuroPi, reason any)

func handlePanicOnScreenLog(e *EuroPi, reason any) {
	if e == nil {
		// can't do anything if it's not enabled
	}

	// force display-logging to enabled
	enableDisplayLogger(e)

	// show the panic on the screen
	log.Panicln(fmt.Sprint(reason))
}

func handlePanicDisplayCrash(e *EuroPi, reason any) {
	if e == nil {
		// can't do anything if it's not enabled
	}

	// display a diagonal line pattern through the screen to show that the EuroPi is crashed
	ymax := int16(output.OLEDHeight) - 1
	for x := -ymax; x < output.OLEDWidth; x += 4 {
		lx, ly := x, int16(0)
		if x < 0 {
			lx = 0
			ly = -x
		}
		e.Display.DrawLine(lx, ly, x+ymax, ymax, output.White)
	}
}
