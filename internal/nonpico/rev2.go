//go:build !pico && (revision2 || europix)
// +build !pico
// +build revision2 europix

package nonpico

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

func init() {
	hardware.SetDetectedRevision(hal.Revision2)
}
