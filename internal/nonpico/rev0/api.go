//go:build !pico
// +build !pico

package rev0

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type voltageOutputMsg struct {
	Kind       string         `json:"kind"`
	HardwareId hal.HardwareId `json:"hardwareId"`
	Voltage    float32        `json:"voltage"`
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
