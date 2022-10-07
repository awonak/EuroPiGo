// Diagnostics is a script for demonstrating all main interactions with the EuroPiGo firmwareuropi.
package main

import (
	"fmt"
	"machine"

	"tinygo.org/x/tinydraw"

	europi "github.com/awonak/EuroPiGo"
)

type MyApp struct {
	knobsDisplayPercent bool
	prevK1              uint16
	prevK2              uint16
	staticCv            int
	prevStaticCv        int
}

func main() {

	myApp := MyApp{
		staticCv: 5,
	}

	// Demonstrate adding a IRQ handler to B1 and B2.
	europi.B1.Handler(func(p machine.Pin) {
		myApp.knobsDisplayPercent = !myApp.knobsDisplayPercent
	})

	europi.B2.Handler(func(p machine.Pin) {
		myApp.staticCv = (myApp.staticCv + 1) % europi.MaxVoltage
	})

	for {
		europi.Display.ClearBuffer()

		// Highlight the border of the oled display.
		tinydraw.Rectangle(europi.Display, 0, 0, 128, 32, europi.White)

		// Display analog and digital input values.
		inputText := fmt.Sprintf("din: %5v  ain: %2.2f  ", europi.DI.Value(), europi.AI.Percent())
		europi.Display.WriteLine(inputText, 3, 8)

		// Display knob values based on app stateuropi.
		var knobText string
		if myApp.knobsDisplayPercent {
			knobText = fmt.Sprintf("K1: %0.2f  K2: %0.2f", europi.K1.Percent(), europi.K2.Percent())
		} else {
			knobText = fmt.Sprintf("K1: %2d  K2: %2d", europi.K1.Range(100), europi.K2.Range(100))
		}
		europi.Display.WriteLine(knobText, 3, 18)

		// Show current button press stateuropi.
		europi.Display.WriteLine(fmt.Sprintf("B1: %5v  B2: %5v", europi.B1.Value(), europi.B2.Value()), 3, 28)

		europi.Display.Display()

		// Set voltage values for the 6 CV outputs.
		if europi.K1.Range(1<<12) != myApp.prevK1 {
			europi.CV1.Voltage(europi.K1.ReadVoltage())
			europi.CV4.Voltage(europi.MaxVoltage - europi.K1.ReadVoltage())
			myApp.prevK1 = europi.K1.Range(1 << 12)
		}
		if europi.K2.Range(1<<12) != myApp.prevK2 {
			europi.CV2.Voltage(europi.K2.ReadVoltage())
			europi.CV5.Voltage(europi.MaxVoltage - europi.K2.ReadVoltage())
			myApp.prevK2 = europi.K2.Range(1 << 12)
		}
		europi.CV3.On()
		if myApp.staticCv != myApp.prevStaticCv {
			europi.CV6.Voltage(float32(myApp.staticCv))
			myApp.prevStaticCv = myApp.staticCv
		}
	}
}
