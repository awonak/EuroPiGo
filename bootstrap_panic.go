package europi

import (
	"fmt"
	"log"

	"github.com/awonak/EuroPiGo/experimental/draw"
	"tinygo.org/x/tinydraw"
)

// DefaultPanicHandler is the default handler for panics
// This will be set by the build flag `onscreenpanic` to `handlePanicOnScreenLog`
// Not setting the build flag will set it to `handlePanicDisplayCrash`
var DefaultPanicHandler func(e *EuroPi, reason any)

func handlePanicOnScreenLog(e *EuroPi, reason any) {
	if e == nil {
		// can't do anything if it's not enabled
		return
	}

	// force display-logging to enabled
	enableDisplayLogger(e)

	// show the panic on the screen
	log.Panicln(fmt.Sprint(reason))

	flushDisplayLogger(e)
}

func handlePanicDisplayCrash(e *EuroPi, reason any) {
	if e == nil {
		// can't do anything if it's not enabled
		return
	}

	// display a diagonal line pattern through the screen to show that the EuroPi is crashed
	disp := e.Display
	width, height := disp.Size()
	ymax := height - 1
	for x := -ymax; x < width; x += 4 {
		lx, ly := x, int16(0)
		if x < 0 {
			lx = 0
			ly = -x
		}
		tinydraw.Line(e.Display, lx, ly, x+ymax, ymax, draw.White)
	}
}
