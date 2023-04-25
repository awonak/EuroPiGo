//go:build !pico
// +build !pico

package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

type voltageOutput struct {
	Kind       string         `json:"kind"`
	HardwareId hal.HardwareId `json:"hardwareId"`
	Voltage    float32        `json:"voltage"`
}

// displayMode = displayModeSeparate (0)
type displayOutput struct {
	Kind       string           `json:"kind"`
	HardwareId hal.HardwareId   `json:"hardwareId"`
	Op         rev1.HwDisplayOp `json:"op"`
	Params     []int16          `json:"params"`
}

// displayMode = displayModeCombined (1)
type displayScreenOuptut struct {
	Kind   string `json:"kind"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Data   []byte `json:"data"`
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
