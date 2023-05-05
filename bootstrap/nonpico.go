//go:build !pico
// +build !pico

package bootstrap

import (
	"context"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/internal/nonpico"
)

func nonPicoActivateWebSocket(ctx context.Context, e europi.Hardware) NonPicoWSActivation {
	nonPicoWSApi := nonpico.ActivateWebSocket(ctx, e.Revision())
	return nonPicoWSApi
}

func nonPicoDeactivateWebSocket(e europi.Hardware, nonPicoWSApi NonPicoWSActivation) {
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
