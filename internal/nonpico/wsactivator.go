//go:build !pico
// +build !pico

package nonpico

import (
	"context"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev0"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev1"
)

type WSActivation interface {
	Shutdown() error
}

func ActivateWebSocket(ctx context.Context, revision hal.Revision) WSActivation {
	switch revision {
	case hal.Revision0:
		return rev0.ActivateWebSocket(ctx)
	case hal.Revision1:
		return rev1.ActivateWebSocket(ctx)
	default:
		return nil
	}
}
