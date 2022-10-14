// Diagnostics is a script for demonstrating all main interactions with the EuroPiGo firmware.
package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/tinydraw"

	europi "github.com/awonak/EuroPiGo"
)

type MyApp struct {
	knobsDisplayPercent bool
	displayShouldUpdate bool
	prevK1              uint16
	prevK2              uint16
	staticCv            int
	prevStaticCv        int
}

func (m *MyApp) updateDisplay() {
	if !m.displayShouldUpdate {
		return
	}
	m.displayShouldUpdate = false

	EuroPi := europi.GetInstance()

	EuroPi.Display.ClearBuffer()

	// Highlight the border of the oled display.
	tinydraw.Rectangle(EuroPi.Display, 0, 0, 128, 32, europi.White)

	// Display analog and digital input values.
	inputText := fmt.Sprintf("din: %5v  ain: %2.2f  ", EuroPi.DI.Value(), EuroPi.AI.Percent())
	EuroPi.Display.WriteLine(inputText, 3, 8)

	// Display knob values based on app stateuropi.
	var knobText string
	if m.knobsDisplayPercent {
		knobText = fmt.Sprintf("K1: %0.2f  K2: %0.2f", EuroPi.K1.Percent(), EuroPi.K2.Percent())
	} else {
		knobText = fmt.Sprintf("K1: %2d  K2: %2d", EuroPi.K1.Range(100), EuroPi.K2.Range(100))
	}
	EuroPi.Display.WriteLine(knobText, 3, 18)

	// Show current button press stateuropi.
	EuroPi.Display.WriteLine(fmt.Sprintf("B1: %5v  B2: %5v", EuroPi.B1.Value(), EuroPi.B2.Value()), 3, 28)

	EuroPi.Display.Display()
}
func main() {

	myApp := MyApp{
		staticCv: 5,
	}
	EuroPi := europi.GetInstance()

	// Demonstrate adding a IRQ handler to B1 and B2.
	EuroPi.B1.Handler(func(p machine.Pin) {
		myApp.knobsDisplayPercent = !myApp.knobsDisplayPercent
		myApp.displayShouldUpdate = true
	})

	EuroPi.B2.Handler(func(p machine.Pin) {
		myApp.staticCv = (myApp.staticCv + 1) % europi.MaxVoltage
		myApp.displayShouldUpdate = true

	})

	go europi.DebugMemoryUsedPerSecond()

	for {

		// Set voltage values for the 6 CV outputs.
		if EuroPi.K1.Range(1<<12) != myApp.prevK1 {
			EuroPi.CV1.Voltage(EuroPi.K1.ReadVoltage())
			EuroPi.CV4.Voltage(europi.MaxVoltage - EuroPi.K1.ReadVoltage())
			myApp.prevK1 = EuroPi.K1.Range(1 << 12)
			myApp.displayShouldUpdate = true
		}
		if EuroPi.K2.Range(1<<12) != myApp.prevK2 {
			EuroPi.CV2.Voltage(EuroPi.K2.ReadVoltage())
			EuroPi.CV5.Voltage(europi.MaxVoltage - EuroPi.K2.ReadVoltage())
			myApp.prevK2 = EuroPi.K2.Range(1 << 12)
			myApp.displayShouldUpdate = true
		}
		EuroPi.CV3.On()
		if myApp.staticCv != myApp.prevStaticCv {
			EuroPi.CV6.Voltage(float32(myApp.staticCv))
			myApp.prevStaticCv = myApp.staticCv
			myApp.displayShouldUpdate = true
		}

		myApp.updateDisplay()
		time.Sleep(10 * time.Millisecond)
	}
}
