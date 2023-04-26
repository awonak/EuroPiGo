//go:build !pico
// +build !pico

package nonpico

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev1"
)

type WSActivation interface {
	Shutdown() error
}

func ActivateWebSocket(revision hal.Revision) WSActivation {
	switch revision {
	case hal.Revision1:
		return rev1.ActivateWebSocket()

	default:
		return nil
	}
}
