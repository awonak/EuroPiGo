//go:build !pico
// +build !pico

package nonpico

import "github.com/awonak/EuroPiGo/hardware/hal"

var detectedRevision hal.Revision

func DetectRevision() hal.Revision {
	return detectedRevision
}
