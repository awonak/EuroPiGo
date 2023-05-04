package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
)

type screenBankEntry struct {
	name       string
	logo       string
	screen     *screenBankEntryDetails
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
	europi.UserInterface[europi.Hardware]
	europi.UserInterfaceButton1[europi.Hardware]
	europi.UserInterfaceButton1Long[europi.Hardware]
	europi.UserInterfaceButton1Ex[europi.Hardware]
	europi.UserInterfaceButton2[europi.Hardware]
	europi.UserInterfaceButton2Ex[europi.Hardware]
}
