//go:build pico
// +build pico

package europi

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/internal/pico"
)

func picoEnsureHardware() {
	rev := pico.DetectRevision()
	hardware.SetDetectedRevision(rev)
}

func init() {
	ensureHardware = picoEnsureHardware
}
