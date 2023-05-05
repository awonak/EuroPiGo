package screenbank

import (
	"fmt"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/bootstrap"
)

type ScreenBankOption func(sb *ScreenBank) error

// WithScreen sets up a new screen in the chain
//  logo is the emoji to use (see https://github.com/tinygo-org/tinyfont/blob/release/notoemoji/NotoEmoji-Regular-12pt.go)
func WithScreen(name string, logo string, screen any) ScreenBankOption {
	return func(sb *ScreenBank) error {
		details, ok := getScreen(screen)
		if !ok {
			return fmt.Errorf("screen %q does not implement a variant of bootstrap.UserInterface", name)
		}
		e := screenBankEntry{
			name:       name,
			logo:       logo,
			screen:     details,
			enabled:    true,
			locked:     true,
			lastUpdate: time.Now(),
		}

		sb.bank = append(sb.bank, e)
		return nil
	}
}

func getScreen(screen any) (details screenBankEntryDetails, ok bool) {
	if s, _ := screen.(bootstrap.UserInterface[europi.Hardware]); s != nil {
		details.screen = s
		details.button1, _ = screen.(bootstrap.UserInterfaceButton1[europi.Hardware])
		details.button1Long, _ = screen.(bootstrap.UserInterfaceButton1Long[europi.Hardware])
		details.button1Ex, _ = screen.(bootstrap.UserInterfaceButton1Ex[europi.Hardware])
		details.button2, _ = screen.(bootstrap.UserInterfaceButton2[europi.Hardware])
		details.button2Ex, _ = screen.(bootstrap.UserInterfaceButton2Ex[europi.Hardware])

		ok = true
		return
	}

	if details, ok = getScreenForHardware[*europi.EuroPiPrototype](screen); ok {
		return
	}

	if details, ok = getScreenForHardware[*europi.EuroPi](screen); ok {
		return
	}

	// TODO: add rev2

	return
}

func getScreenForHardware[THardware europi.Hardware](screen any) (details screenBankEntryDetails, ok bool) {
	s, _ := screen.(bootstrap.UserInterface[THardware])
	if s == nil {
		return
	}

	wrapper := &screenHardwareWrapper[THardware]{
		screen: s,
	}

	details.screen = wrapper

	if wrapper.button1, _ = screen.(bootstrap.UserInterfaceButton1[THardware]); wrapper.button1 != nil {
		details.button1 = wrapper
	}
	if wrapper.button1Long, _ = screen.(bootstrap.UserInterfaceButton1Long[THardware]); wrapper.button1Long != nil {
		details.button1Long = wrapper
	}
	if wrapper.button1Ex, _ = screen.(bootstrap.UserInterfaceButton1Ex[THardware]); wrapper.button1Ex != nil {
		details.button1Ex = wrapper
	}
	if wrapper.button2, _ = screen.(bootstrap.UserInterfaceButton2[THardware]); wrapper.button2 != nil {
		details.button2 = wrapper
	}
	if wrapper.button2Ex, _ = screen.(bootstrap.UserInterfaceButton2Ex[THardware]); wrapper.button2Ex != nil {
		details.button2Ex = wrapper
	}

	ok = true
	return
}

type screenHardwareWrapper[THardware europi.Hardware] struct {
	screen      bootstrap.UserInterface[THardware]
	button1     bootstrap.UserInterfaceButton1[THardware]
	button1Ex   bootstrap.UserInterfaceButton1Ex[THardware]
	button1Long bootstrap.UserInterfaceButton1Long[THardware]
	button2     bootstrap.UserInterfaceButton2[THardware]
	button2Ex   bootstrap.UserInterfaceButton2Ex[THardware]
}

func (w *screenHardwareWrapper[THardware]) Start(e europi.Hardware) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.screen.Start(pi)
}

func (w *screenHardwareWrapper[THardware]) Paint(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.screen.Paint(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button1(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button1.Button1(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button1Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button1Ex.Button1Ex(pi, value, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button1Long(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button1Long.Button1Long(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button2(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button2.Button2(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button2Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic("incorrect hardware type conversion")
	}
	w.button2Ex.Button2Ex(pi, value, deltaTime)
}
