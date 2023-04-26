//go:build !pico && (revision1 || europi)
// +build !pico
// +build revision1 europi

package nonpico

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

func init() {
	hardware.SetDetectedRevision(hal.Revision1)
}
