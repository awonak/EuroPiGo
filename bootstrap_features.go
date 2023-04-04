package europi

import (
	"log"
	"os"

	"github.com/heucuva/europi/experimental/displaylogger"
)

func enableDisplayLogger(e *EuroPi) {
	log.SetFlags(0)
	log.SetOutput(&displaylogger.Logger{
		Display: e.Display,
	})
}

func disableDisplayLogger(e *EuroPi) {
	log.SetOutput(os.Stdout)
}
