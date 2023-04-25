//go:build !pico && revision1
// +build !pico,revision1

package europi

import (
	"context"
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/internal/nonpico/events"
	"github.com/awonak/EuroPiGo/internal/nonpico/ws"
)

//go:embed internal/nonpico/site
var nonpicoSiteContent embed.FS

func nonPicoActivateWebSocket(e *EuroPi) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()

		subFS, _ := fs.Sub(nonpicoSiteContent, "internal/nonpico/site")
		http.Handle("/", http.FileServer(http.FS(subFS)))
		http.HandleFunc("/ws", nonPicoApiHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	return ctx, cancel
}

func nonPicoDeactivateWebSocket(e *EuroPi, cancel context.CancelFunc) {
	cancel()
}

func nonPicoApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, "nonPicoApiHandler")

	sock, err := ws.Upgrade(w, r)
	if err != nil {
		log.Println("failed to upgrade websocket connection:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer sock.Close()

	type voltageOutput struct {
		Kind       string         `json:"kind"`
		HardwareId hal.HardwareId `json:"hardwareId"`
		Voltage    float32        `json:"voltage"`
	}
	events.SetupVoltageOutputListeners(func(id hal.HardwareId, voltage float32) {
		_ = sock.WriteJSON(voltageOutput{
			Kind:       "voltageOutput",
			HardwareId: id,
			Voltage:    voltage,
		})
	})

	type displayOutput struct {
		Kind       string           `json:"kind"`
		HardwareId hal.HardwareId   `json:"hardwareId"`
		Op         rev1.HwDisplayOp `json:"op"`
		Params     []int16          `json:"params"`
	}

	events.SetupDisplayOutputListener(func(id hal.HardwareId, op rev1.HwDisplayOp, params []int16) {
		_ = sock.WriteJSON(displayOutput{
			Kind:       "displayOutput",
			HardwareId: id,
			Op:         op,
			Params:     params,
		})
	})

	type kind struct {
		Kind string `json:"kind"`
	}

	type setDigitalInput struct {
		Kind       string         `json:"kind"`
		HardwareId hal.HardwareId `json:"hardwareId"`
		Value      bool           `json:"value"`
	}

	type setAnalogInput struct {
		Kind       string         `json:"kind"`
		HardwareId hal.HardwareId `json:"hardwareId"`
		Voltage    float32        `json:"voltage"`
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
			var di setDigitalInput
			if err := json.Unmarshal(blob, &di); err != nil {
				sock.SetError(err)
				break
			}
			events.SetDigitalInput(di.HardwareId, di.Value)

		case "setAnalogInput":
			var ai setAnalogInput
			if err := json.Unmarshal(blob, &ai); err != nil {
				sock.SetError(err)
				break
			}
			events.SetAnalogInput(ai.HardwareId, ai.Voltage)

		default:
			// ignore
		}
	}
}

func init() {
	activateNonPicoWebSocket = nonPicoActivateWebSocket
	deactivateNonPicoWebSocket = nonPicoDeactivateWebSocket
}
