// Diagnostics is a script for demonstrating all main interactions with the europi-go firmware.
package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/tinydraw"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/input"
	"github.com/heucuva/europi/output"
)

type MyApp struct {
	knobsDisplayPercent bool
	prevK1              uint16
	prevK2              uint16
	staticCv            int
	prevStaticCv        int
}

var myApp MyApp

func startLoop(e *europi.EuroPi) {
	myApp.staticCv = 5

	// Demonstrate adding a IRQ handler to B1 and B2.
	e.B1.Handler(func(p machine.Pin) {
		myApp.knobsDisplayPercent = !myApp.knobsDisplayPercent
	})

	e.B2.Handler(func(p machine.Pin) {
		myApp.staticCv = (myApp.staticCv + 1) % input.MaxVoltage
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	e.Display.ClearBuffer()

	// Highlight the border of the oled display.
	tinydraw.Rectangle(e.Display, 0, 0, 128, 32, output.White)

	// Display analog and digital input values.
	inputText := fmt.Sprintf("din: %5v  ain: %2.2f  ", e.DI.Value(), e.AI.Percent())
	e.Display.WriteLine(inputText, 3, 8)

	// Display knob values based on app state.
	var knobText string
	if myApp.knobsDisplayPercent {
		knobText = fmt.Sprintf("K1: %0.2f  K2: %0.2f", e.K1.Percent(), e.K2.Percent())
	} else {
		knobText = fmt.Sprintf("K1: %2d  K2: %2d", e.K1.Range(100), e.K2.Range(100))
	}
	e.Display.WriteLine(knobText, 3, 18)

	// Show current button press state.
	e.Display.WriteLine(fmt.Sprintf("B1: %5v  B2: %5v", e.B1.Value(), e.B2.Value()), 3, 28)

	e.Display.Display()

	// Set voltage values for the 6 CV outputs.
	if e.K1.Range(1<<12) != myApp.prevK1 {
		e.CV1.SetVoltage(e.K1.ReadVoltage())
		e.CV4.SetVoltage(output.MaxVoltage - e.K1.ReadVoltage())
		myApp.prevK1 = e.K1.Range(1 << 12)
	}
	if e.K2.Range(1<<12) != myApp.prevK2 {
		e.CV2.SetVoltage(e.K2.ReadVoltage())
		e.CV5.SetVoltage(output.MaxVoltage - e.K2.ReadVoltage())
		myApp.prevK2 = e.K2.Range(1 << 12)
	}
	e.CV3.On()
	if myApp.staticCv != myApp.prevStaticCv {
		e.CV6.SetVoltage(float32(myApp.staticCv))
		myApp.prevStaticCv = myApp.staticCv
	}
}

func main() {
	europi.Bootstrap(
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
	)
}
