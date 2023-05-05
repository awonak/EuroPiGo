package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/bootstrap"
)

type entryWrapper[THardware europi.Hardware] struct {
	screen      bootstrap.UserInterface[THardware]
	button1     bootstrap.UserInterfaceButton1[THardware]
	button1Long bootstrap.UserInterfaceButton1Long[THardware]
	button1Ex   bootstrap.UserInterfaceButton1Ex[THardware]
	button2     bootstrap.UserInterfaceButton2[THardware]
	button2Ex   bootstrap.UserInterfaceButton2Ex[THardware]
}

func (w *entryWrapper[THardware]) Start(e europi.Hardware) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.screen.Start(pi)
}

func (w *entryWrapper[THardware]) Paint(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.screen.Paint(pi, deltaTime)
}

func (w *entryWrapper[THardware]) Button1(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button1.Button1(pi, deltaTime)
}

func (w *entryWrapper[THardware]) Button1Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button1Ex.Button1Ex(pi, value, deltaTime)
}

func (w *entryWrapper[THardware]) Button1Long(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button1Long.Button1Long(pi, deltaTime)
}

func (w *entryWrapper[THardware]) Button2(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button2.Button2(pi, deltaTime)
}

func (w *entryWrapper[THardware]) Button2Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button2Ex.Button2Ex(pi, value, deltaTime)
}
