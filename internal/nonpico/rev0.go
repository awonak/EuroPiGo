//go:build !pico && (revision0 || europiproto || europiprototype)
// +build !pico
// +build revision0 europiproto europiprototype

package nonpico

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

func init() {
	detectedRevision = hal.Revision0
}
