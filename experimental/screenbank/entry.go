package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
)

type entry struct {
	name       string
	logo       string
	screen     entryWrapper[europi.Hardware]
	enabled    bool
	locked     bool
	lastUpdate time.Time
}

func (e *entry) lock() {
	if e.locked {
		return
	}

	e.locked = true
}

func (e *entry) unlock() {
	if !e.enabled {
		return
	}

	e.locked = false
}
