//go:build pico
// +build pico

package europi

import (
	"sync"

	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/internal/pico"
)

var picoEnsureHardwareOnce sync.Once

func picoEnsureHardware() {
	picoEnsureHardwareOnce.Do(func() {
		rev := pico.DetectRevision()
		hardware.SetDetectedRevision(rev)
	})
}

func init() {
	ensureHardware = picoEnsureHardware
}
