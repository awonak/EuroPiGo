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
		e := entry{
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

func getScreen(screen any) (details entryWrapper[europi.Hardware], ok bool) {
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

func getScreenForHardware[THardware europi.Hardware](screen any) (details entryWrapper[europi.Hardware], ok bool) {
	s, _ := screen.(bootstrap.UserInterface[THardware])
	if s == nil {
		return
	}

	wrapper := &entryWrapper[THardware]{
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
