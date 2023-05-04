//go:build !pico
// +build !pico

package europi

import (
	"context"

	"github.com/awonak/EuroPiGo/internal/nonpico"
)

func nonPicoActivateWebSocket(ctx context.Context, e Hardware) NonPicoWSActivation {
	nonPicoWSApi := nonpico.ActivateWebSocket(ctx, e.Revision())
	return nonPicoWSApi
}

func nonPicoDeactivateWebSocket(e Hardware, nonPicoWSApi NonPicoWSActivation) {
	if nonPicoWSApi != nil {
		if err := nonPicoWSApi.Shutdown(); err != nil {
			panic(err)
		}
	}
}

func init() {
	activateNonPicoWebSocket = nonPicoActivateWebSocket
	deactivateNonPicoWebSocket = nonPicoDeactivateWebSocket
}
