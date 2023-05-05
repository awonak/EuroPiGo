//go:build !pico
// +build !pico

package europi

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/internal/nonpico"
)

func nonPicoEnsureHardware() {
	rev := nonpico.DetectRevision()
	hardware.SetDetectedRevision(rev)
}

func init() {
	ensureHardware = nonPicoEnsureHardware
}
