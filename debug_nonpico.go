//go:build !pico
// +build !pico

package bootstrap

import (
	"context"

	"github.com/awonak/EuroPiGo/internal/nonpico"
)

func nonPicoActivateWebSocket(ctx context.Context, e Hardware) NonPicoWSActivation {
	nonPicoWSApi := nonpico.ActivateWebSocket(ctx, e.Revision())
	return nonPicoWSApi
}

func init() {
	activateNonPicoWebSocket = nonPicoActivateWebSocket
}
