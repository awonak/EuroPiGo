package input

import (
	"machine"
)

// AnalogReader is an interface for common analog read methods for knobs and cv input.
type AnalogReader interface {
	Samples(samples uint16)
	ReadVoltage() float32
	Percent() float32
	Range(steps uint16) uint16
}

func init() {
	machine.InitADC()
}
