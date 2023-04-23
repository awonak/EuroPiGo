package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
)

type screenBankEntry struct {
	name       string
	logo       string
	screen     europi.UserInterface
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
