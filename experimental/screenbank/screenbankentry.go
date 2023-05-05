package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/bootstrap"
)

type screenBankEntry struct {
	name       string
	logo       string
	screen     screenBankEntryDetails
	enabled    bool
	locked     bool
	lastUpdate time.Time
}

func (e *screenBankEntry) lock() {
	if e.locked {
		return
	}

	e.locked = true
}

func (e *screenBankEntry) unlock() {
	if !e.enabled {
		return
	}

	e.locked = false
}

type screenBankEntryDetails struct {
	screen      bootstrap.UserInterface[europi.Hardware]
	button1     bootstrap.UserInterfaceButton1[europi.Hardware]
	button1Long bootstrap.UserInterfaceButton1Long[europi.Hardware]
	button1Ex   bootstrap.UserInterfaceButton1Ex[europi.Hardware]
	button2     bootstrap.UserInterfaceButton2[europi.Hardware]
	button2Ex   bootstrap.UserInterfaceButton2Ex[europi.Hardware]
}
