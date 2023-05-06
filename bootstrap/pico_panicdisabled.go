//go:build pico && !onscreenpanic
// +build pico,!onscreenpanic

package bootstrap

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

func init() {
	hardware.OnRevisionDetected(func(revision hal.Revision) {
		switch revision {
		case hal.RevisionUnknown, hal.EuroPiProto:
			DefaultPanicHandler = handlePanicLogger
		default:
			DefaultPanicHandler = handlePanicDisplayCrash
		}
	})
}
