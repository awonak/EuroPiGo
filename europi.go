package europi // import "github.com/awonak/EuroPiGo"

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	_ "github.com/awonak/EuroPiGo/internal/nonpico"
	_ "github.com/awonak/EuroPiGo/internal/pico"
)

type (
	// Hardware is the collection of component wrappers used to interact with the module.
	Hardware = hal.Hardware

	// EuroPiPrototype is the revision 0 hardware
	EuroPiPrototype = rev0.EuroPiPrototype
	// EuroPi is the revision 1 hardware
	EuroPi = rev1.EuroPi
	// TODO: add rev2
)

// New will return a new EuroPi struct based on the detected hardware revision
func New() Hardware {
	// ensure our hardware has been identified
	ensureHardware()

	// blocks until revision has been identified
	revision := hardware.GetRevision()
	return NewFrom(revision)
}

// NewFrom will return a new EuroPi struct based on a specific revision
func NewFrom(revision hal.Revision) Hardware {
	if revision == hal.RevisionUnknown {
		// unknown revision
		return nil
	}

	// ensure our hardware has been identified
	ensureHardware()

	// this will block until the hardware components are initialized
	hardware.WaitForReady()

	switch revision {
	case hal.Revision0:
		return rev0.Pi
	case hal.Revision1:
		return rev1.Pi
	case hal.Revision2:
		// TODO: add rev2
		return nil
	default:
		return nil
	}
}

// Display returns the primary display from the hardware interface, if it has one
func Display(e Hardware) hal.DisplayOutput {
	if e == nil {
		return nil
	}

	switch e.Revision() {
	case hal.Revision1:
		return e.(*rev1.EuroPi).OLED
	case hal.Revision2:
		// TODO: add rev2
		//return e.(*rev2.EuroPiX).Display
	}
	return nil
}

// Button returns a button input at the specified index from the hardware interface,
// if it has one there
func Button(e Hardware, idx int) hal.ButtonInput {
	if e == nil {
		return nil
	}

	switch e.Revision() {
	case hal.Revision0:
		return e.(*rev0.EuroPiPrototype).Button(idx)
	case hal.Revision1:
		return e.(*rev1.EuroPi).Button(idx)
	case hal.Revision2:
		// TODO: add rev2
		//return e.(*rev2.EuroPiX).Button(idx)
	}
	return nil
}

// Knob returns a knob input at the specified index from the hardware interface,
// if it has one there
func Knob(e Hardware, idx int) hal.KnobInput {
	if e == nil {
		return nil
	}

	switch e.Revision() {
	case hal.Revision0:
		return e.(*rev0.EuroPiPrototype).Knob(idx)
	case hal.Revision1:
		return e.(*rev1.EuroPi).Knob(idx)
	case hal.Revision2:
		// TODO: add rev2
		//return e.(*rev2.EuroPiX).Knob(idx)
	}
	return nil
}

// ensureHardware is part of the hardware setup system.
// It will be set to a function that can properly configure the hardware for use.
var ensureHardware func()
