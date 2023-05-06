package screenbank

import (
	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/bootstrap"
)

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
