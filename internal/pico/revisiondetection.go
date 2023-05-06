//go:build pico
// +build pico

package pico

import (
	"machine"
	"runtime/interrupt"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

func configureAndGetPinLevel(pin machine.Pin) bool {
	pin.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})
	state := interrupt.Disable()
	level := pin.Get()
	interrupt.Restore(state)
	return level
}

func GetDetectedRevisionBits() int {
	var revision int
	if gp6 := configureAndGetPinLevel(machine.GPIO6); gp6 {
		revision |= 0b0001
	}
	if gp7 := configureAndGetPinLevel(machine.GPIO7); gp7 {
		revision |= 0b0010
	}
	if gp8 := configureAndGetPinLevel(machine.GPIO8); gp8 {
		revision |= 0b0100
	}
	if gp9 := configureAndGetPinLevel(machine.GPIO9); gp9 {
		revision |= 0b1000
	}
	return revision
}

func DetectRevision() hal.Revision {
	revBits := GetDetectedRevisionBits()
	switch revBits {
	case 0: // 0000
		if rev1AsRev0 {
			return hal.Revision0
		}
		return hal.Revision1
	case 1: // 0001
		return hal.Revision2
	default: // not yet known or maybe Revision0 / EuroPi-Proto?
		return hal.RevisionUnknown
	}
}

var rev1AsRev0 bool
