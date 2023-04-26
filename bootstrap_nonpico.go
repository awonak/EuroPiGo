//go:build !pico
// +build !pico

package europi

import "github.com/awonak/EuroPiGo/internal/nonpico"

func nonPicoActivateWebSocket(e *EuroPi) nonPicoWSActivation {
	nonPicoWSApi := nonpico.ActivateWebSocket(e.Revision)
	return nonPicoWSApi
}

func nonPicoDeactivateWebSocket(e *EuroPi, nonPicoWSApi nonPicoWSActivation) {
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
