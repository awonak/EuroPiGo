//go:build !pico
// +build !pico

package nonpico

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev1"
)

func initRevision0() {
	rev0.DoInit()
}

func initRevision1() {
	rev1.DoInit()
}

func initRevision2() {
	//TODO: rev2.DoInit()
}

func init() {
	hardware.OnRevisionDetected <- func(revision hal.Revision) {
		switch revision {
		case hal.Revision0:
			initRevision0()
		case hal.Revision1:
			initRevision1()
		case hal.Revision2:
			initRevision2()
		default:
		}
		hardware.SetReady()
	}
}
