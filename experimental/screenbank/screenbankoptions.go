package screenbank

import (
	"fmt"
	"time"

	europi "github.com/awonak/EuroPiGo"
)

type ScreenBankOption func(sb *ScreenBank) error

// WithScreen sets up a new screen in the chain
//  logo is the emoji to use (see https://github.com/tinygo-org/tinyfont/blob/release/notoemoji/NotoEmoji-Regular-12pt.go)
func WithScreen(name string, logo string, screen any) ScreenBankOption {
	return func(sb *ScreenBank) error {
		details := getScreen(screen)
		if details == nil {
			return fmt.Errorf("screen %q does not implement a variant of europi.UserInterface", name)
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

func getScreen(screen any) *screenBankEntryDetails {
	if s, _ := screen.(europi.UserInterface[europi.Hardware]); s != nil {
		details := &screenBankEntryDetails{
			UserInterface: s,
		}

		details.UserInterfaceButton1, _ = screen.(europi.UserInterfaceButton1[europi.Hardware])
		details.UserInterfaceButton1Long, _ = screen.(europi.UserInterfaceButton1Long[europi.Hardware])
		details.UserInterfaceButton1Ex, _ = screen.(europi.UserInterfaceButton1Ex[europi.Hardware])
		details.UserInterfaceButton2, _ = screen.(europi.UserInterfaceButton2[europi.Hardware])
		details.UserInterfaceButton2Ex, _ = screen.(europi.UserInterfaceButton2Ex[europi.Hardware])

		return details
	}

	if s := getScreenForHardware[*europi.EuroPiPrototype](screen); s != nil {
		return s
	}

	if s := getScreenForHardware[*europi.EuroPi](screen); s != nil {
		return s
	}

	// TODO: add rev2

	return nil
}

func getScreenForHardware[THardware europi.Hardware](screen any) *screenBankEntryDetails {
	s, _ := screen.(europi.UserInterface[THardware])
	if s == nil {
		return nil
	}

	wrapper := &screenHardwareWrapper[THardware]{
		UserInterface: s,
	}

	details := &screenBankEntryDetails{
		UserInterface: wrapper,
	}

	if wrapper.button1, _ = screen.(europi.UserInterfaceButton1[THardware]); wrapper.button1 != nil {
		details.UserInterfaceButton1 = wrapper
	}
	if wrapper.button1Long, _ = screen.(europi.UserInterfaceButton1Long[THardware]); wrapper.button1Long != nil {
		details.UserInterfaceButton1Long = wrapper
	}
	if wrapper.button1Ex, _ = screen.(europi.UserInterfaceButton1Ex[THardware]); wrapper.button1Ex != nil {
		details.UserInterfaceButton1Ex = wrapper
	}
	if wrapper.button2, _ = screen.(europi.UserInterfaceButton2[THardware]); wrapper.button2 != nil {
		details.UserInterfaceButton2 = wrapper
	}
	if wrapper.button2Ex, _ = screen.(europi.UserInterfaceButton2Ex[THardware]); wrapper.button2Ex != nil {
		details.UserInterfaceButton2Ex = wrapper
	}

	return details
}

type screenHardwareWrapper[THardware europi.Hardware] struct {
	europi.UserInterface[THardware]
	button1     europi.UserInterfaceButton1[THardware]
	button1Ex   europi.UserInterfaceButton1Ex[THardware]
	button1Long europi.UserInterfaceButton1Long[THardware]
	button2     europi.UserInterfaceButton2[THardware]
	button2Ex   europi.UserInterfaceButton2Ex[THardware]
}

func (w *screenHardwareWrapper[THardware]) Start(e europi.Hardware) {
	pi, _ := e.(THardware)
	w.UserInterface.Start(pi)
}

func (w *screenHardwareWrapper[THardware]) Paint(e europi.Hardware, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	w.UserInterface.Paint(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button1(e europi.Hardware, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	w.button1.Button1(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button1Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	w.button1Ex.Button1Ex(pi, value, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button1Long(e europi.Hardware, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	w.button1Long.Button1Long(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button2(e europi.Hardware, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	w.button2.Button2(pi, deltaTime)
}

func (w *screenHardwareWrapper[THardware]) Button2Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	w.button2Ex.Button2Ex(pi, value, deltaTime)
}

//	Button1Debounce() time.Duration
//	Button2Debounce() time.Duration
//	Button2Long(e THardware, deltaTime time.Duration)
