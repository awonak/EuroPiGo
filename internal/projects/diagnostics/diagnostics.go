// Diagnostics is a script for demonstrating all main interactions with the europi-go firmware.
package main

import (
	"fmt"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont/proggy"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
)

type MyApp struct {
	knobsDisplayPercent bool
	prevK1              uint16
	prevK2              uint16
	staticCv            int
	prevStaticCv        int
}

var myApp MyApp

func appStart(e *europi.EuroPi) {
	myApp.staticCv = 5

	// Demonstrate adding a IRQ handler to B1 and B2.
	e.B1.Handler(func(_ bool, _ time.Duration) {
		myApp.knobsDisplayPercent = !myApp.knobsDisplayPercent
	})

	e.B2.Handler(func(_ bool, _ time.Duration) {
		myApp.staticCv = (myApp.staticCv + 1) % int(e.K1.MaxVoltage())
	})
}

var (
	DefaultFont = &proggy.TinySZ8pt7b
)

func mainLoop(e *europi.EuroPi) {
	e.OLED.ClearBuffer()

	// Highlight the border of the oled display.
	_ = tinydraw.Rectangle(e.OLED, 0, 0, 128, 32, draw.White)

	writer := fontwriter.Writer{
		Display: e.OLED,
		Font:    DefaultFont,
	}

	// Display analog and digital input values.
	inputText := fmt.Sprintf("din: %5v  ain: %2.2f  ", e.DI.Value(), e.AI.Percent())
	writer.WriteLine(inputText, 3, 8, draw.White)

	// Display knob values based on app state.
	var knobText string
	if myApp.knobsDisplayPercent {
		knobText = fmt.Sprintf("K1: %0.2f  K2: %0.2f", e.K1.Percent(), e.K2.Percent())
	} else {
		knobText = fmt.Sprintf("K1: %3d K2: %3d", int(e.K1.Percent()*100), int(e.K2.Percent()*100))
	}
	writer.WriteLine(knobText, 3, 18, draw.White)

	// Show current button press state.
	writer.WriteLine(fmt.Sprintf("B1: %5v  B2: %5v", e.B1.Value(), e.B2.Value()), 3, 28, draw.White)

	_ = e.OLED.Display()

	// Set voltage values for the 6 CV outputs.
	if kv := uint16(e.K1.Percent() * float32(1<<12)); kv != myApp.prevK1 {
		e.CV1.SetVoltage(e.K1.ReadVoltage())
		e.CV4.SetVoltage(e.CV4.MaxVoltage() - e.K1.ReadVoltage())
		myApp.prevK1 = kv
	}
	if kv := uint16(e.K2.Percent() * float32(1<<12)); kv != myApp.prevK2 {
		e.CV2.SetVoltage(e.K2.ReadVoltage())
		e.CV5.SetVoltage(e.CV5.MaxVoltage() - e.K2.ReadVoltage())
		myApp.prevK2 = kv
	}
	e.CV3.SetCV(1.0)
	if myApp.staticCv != myApp.prevStaticCv {
		e.CV6.SetVoltage(float32(myApp.staticCv))
		myApp.prevStaticCv = myApp.staticCv
	}
}

func main() {
	e, _ := europi.New().(*europi.EuroPi)
	if e == nil {
		panic("europi not detected")
	}

	appStart(e)
	for {
		mainLoop(e)
		time.Sleep(time.Millisecond)
	}
}
