//go:build !pico
// +build !pico

package rev0

import (
	"context"
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/internal/nonpico/ws"
)

type WSActivation struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func ActivateWebSocket(ctx context.Context) *WSActivation {
	a := &WSActivation{}

	a.Start(ctx)

	return a
}

func (a *WSActivation) Shutdown() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

//go:embed site
var nonPicoSiteContent embed.FS

func (a *WSActivation) Start(ctx context.Context) {
	a.ctx, a.cancel = context.WithCancel(ctx)

	// initialize default state
	setupDefaultState()

	go func() {
		defer a.cancel()

		subFS, _ := fs.Sub(nonPicoSiteContent, "site")
		http.Handle("/", http.FileServer(http.FS(subFS)))
		http.HandleFunc("/ws", a.apiHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()
}

func (a *WSActivation) apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, "rev0.apiHandler")

	if r.Body != nil {
		// just in case someone sent us a body
		defer r.Body.Close()
	}

	sock, err := ws.Upgrade(w, r)
	if err != nil {
		log.Println("failed to upgrade websocket connection:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer sock.Close()

	setupVoltageOutputListeners(func(id hal.HardwareId, voltage float32) {
		_ = sock.WriteJSON(voltageOutputMsg{
			Kind:       "voltageOutput",
			HardwareId: id,
			Voltage:    voltage,
		})
	})

	type kind struct {
		Kind string `json:"kind"`
	}

	for {
		// test for doneness
		select {
		case <-sock.Done():
			break
		default:
		}

		blob, err := sock.ReadMessage()
		if err != nil {
			break
		}

		var k kind
		if err := json.Unmarshal(blob, &k); err != nil {
			sock.SetError(err)
			break
		}

		switch k.Kind {
		case "setDigitalInput":
			var di setDigitalInputMsg
			if err := json.Unmarshal(blob, &di); err != nil {
				sock.SetError(err)
				break
			}
			setDigitalInput(di.HardwareId, di.Value)

		case "setAnalogInput":
			var ai setAnalogInputMsg
			if err := json.Unmarshal(blob, &ai); err != nil {
				sock.SetError(err)
				break
			}
			setAnalogInput(ai.HardwareId, ai.Voltage)

		default:
			// ignore
		}
	}
}
