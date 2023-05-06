//go:build !pico
// +build !pico

package europi

import (
	"sync"

	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/internal/nonpico"
)

var nonPicoEnsureHardwareOnce sync.Once

func nonPicoEnsureHardware() {
	nonPicoEnsureHardwareOnce.Do(func() {
		rev := nonpico.DetectRevision()
		hardware.SetDetectedRevision(rev)
	})
}

func init() {
	ensureHardware = nonPicoEnsureHardware
}
