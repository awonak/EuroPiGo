//go:build !pico
// +build !pico

package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type voltageOutputMsg struct {
	Kind       string         `json:"kind"`
	HardwareId hal.HardwareId `json:"hardwareId"`
	Voltage    float32        `json:"voltage"`
}

// displayMode = displayModeSeparate (0)
type displayOutputMsg struct {
	Kind       string         `json:"kind"`
	HardwareId hal.HardwareId `json:"hardwareId"`
	Op         HwDisplayOp    `json:"op"`
	Params     []int16        `json:"params"`
}

// displayMode = displayModeCombined (1)
type displayScreenOuptutMsg struct {
	Kind   string `json:"kind"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Data   []byte `json:"data"`
}

type setDigitalInputMsg struct {
	Kind       string         `json:"kind"`
	HardwareId hal.HardwareId `json:"hardwareId"`
	Value      bool           `json:"value"`
}

type setAnalogInputMsg struct {
	Kind       string         `json:"kind"`
	HardwareId hal.HardwareId `json:"hardwareId"`
	Voltage    float32        `json:"voltage"`
}
