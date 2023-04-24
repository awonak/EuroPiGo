//go:build !onscreenpanic
// +build !onscreenpanic

package europi

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

func init() {
	switch hardware.RevisionDetection() {
	case hal.RevisionUnknown:
		DefaultPanicHandler = handlePanicLogger
	default:
		DefaultPanicHandler = handlePanicDisplayCrash
	}
}
