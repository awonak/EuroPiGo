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
	"strconv"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/internal/nonpico/rev1/events"
	"github.com/awonak/EuroPiGo/internal/nonpico/ws"
)

type WSActivation struct {
	ctx         context.Context
	cancel      context.CancelFunc
	displayMode displayMode
}

func ActivateWebSocket() *WSActivation {
	a := &WSActivation{}

	a.Start(context.Background())

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
	log.Println(r.URL, "rev1.apiHandler")

	if r.Body != nil {
		// just in case someone sent us a body
		defer r.Body.Close()
	}

	q := r.URL.Query()
	dm, _ := strconv.Atoi(q.Get("displayMode"))
	a.displayMode = displayMode(dm)

	sock, err := ws.Upgrade(w, r)
	if err != nil {
		log.Println("failed to upgrade websocket connection:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer sock.Close()

	events.SetupVoltageOutputListeners(func(id hal.HardwareId, voltage float32) {
		_ = sock.WriteJSON(voltageOutput{
			Kind:       "voltageOutput",
			HardwareId: id,
			Voltage:    voltage,
		})
	})

	displayWidth, displayHeight := 128, 32
	displayScreenOutputMsg := displayScreenOuptut{
		Kind:   "displayScreenOutput",
		Width:  displayWidth,
		Height: displayHeight,
		Data:   make([]byte, displayWidth*displayHeight*4),
	}
	events.SetupDisplayOutputListener(func(id hal.HardwareId, op rev1.HwDisplayOp, params []int16) {
		switch a.displayMode {
		case displayModeCombined:
			switch op {
			case rev1.HwDisplayOpClearBuffer:
				for i := range displayScreenOutputMsg.Data {
					displayScreenOutputMsg.Data[i] = 0
				}
			case rev1.HwDisplayOpSetPixel:
				y, x := int(params[1]), int(params[0])
				if y < 0 || y >= displayHeight || x < 0 || x >= displayWidth {
					break
				}
				pos := (int(params[1])*displayWidth + int(params[0])) * 4
				displayScreenOutputMsg.Data[pos] = byte(params[2])
				displayScreenOutputMsg.Data[pos+1] = byte(params[3])
				displayScreenOutputMsg.Data[pos+2] = byte(params[4])
				displayScreenOutputMsg.Data[pos+3] = byte(params[5])
			case rev1.HwDisplayOpDisplay:
				_ = sock.WriteJSON(displayScreenOutputMsg)
			default:
			}

		default:
			_ = sock.WriteJSON(displayOutput{
				Kind:       "displayOutput",
				HardwareId: id,
				Op:         op,
				Params:     params,
			})
		}
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
