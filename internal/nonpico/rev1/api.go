//go:build !pico
// +build !pico

package rev1

import (
	"context"
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev1/events"
	"github.com/awonak/EuroPiGo/internal/nonpico/ws"
)

//go:embed site
var nonPicoSiteContent embed.FS

type WSActivation struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (a *WSActivation) Shutdown() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

func ActivateWebSocket() *WSActivation {
	ctx, cancel := context.WithCancel(context.Background())

	a := &WSActivation{
		ctx:    ctx,
		cancel: cancel,
	}

	go func() {
		defer cancel()

		subFS, _ := fs.Sub(nonPicoSiteContent, "site")
		http.Handle("/", http.FileServer(http.FS(subFS)))
		http.HandleFunc("/ws", apiHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	return a
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, "rev1.apiHandler")

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
