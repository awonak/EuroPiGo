package main

import (
	"fmt"
	"machine"

	"tinygo.org/x/tinydraw"

	europi "github.com/awonak/EuroPiGo"
)

type MyApp struct {
	knobsDisplayPercent bool
	staticCv            int
	prevStaticCv        int
	prevK1              int
	prevK2              int
}

func main() {

	myApp := MyApp{
		staticCv: 5,
	}

	e := europi.New()

	// Demonstrate adding a IRQ handler to B1 and B2.
	e.B1.Handler(func(p machine.Pin) {
		myApp.knobsDisplayPercent = !myApp.knobsDisplayPercent
	})

	e.B2.Handler(func(p machine.Pin) {
		myApp.staticCv = (myApp.staticCv + 1) % europi.MaxVoltage
	})

	for {
		e.Display.ClearBuffer()

		// Highlight the border of the oled display.
		tinydraw.Rectangle(e.Display, 0, 0, 128, 32, europi.White)

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
		if e.K1.Range(1<<10) != myApp.prevK1 {
			e.CV1.Voltage(e.K1.ReadVoltage())
			e.CV4.Voltage(europi.MaxVoltage - e.K1.ReadVoltage())
			myApp.prevK1 = e.K1.Range(1 << 10)
		}
		if e.K2.Range(1<<10) != myApp.prevK2 {
			e.CV2.Voltage(e.K2.ReadVoltage())
			e.CV5.Voltage(europi.MaxVoltage - e.K2.ReadVoltage())
			myApp.prevK2 = e.K2.Range(1 << 10)
		}
		e.CV3.On()
		if myApp.staticCv != myApp.prevStaticCv {
			e.CV6.Voltage(float32(myApp.staticCv))
			myApp.prevStaticCv = myApp.staticCv
		}
	}
}
