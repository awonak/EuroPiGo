package common

import "github.com/awonak/EuroPiGo/hardware/hal"

// HwMessageDigitalValue represents a digital value update
type HwMessageDigitalValue struct {
	Value bool
}

// HwMessageADCValue represents an ADC value update
type HwMessageADCValue struct {
	Value uint16
}

// HwMessageInterrupt represents an interrupt
type HwMessageInterrupt struct {
	Change hal.ChangeFlags
}

// HwMessagePwmValue represents a pulse width modulator value update
type HwMessagePwmValue struct {
	Value   uint16
	Voltage float32
}

// HwMessageDisplay represents a display update.
type HwMessageDisplay struct {
	Op       HwDisplayOp
	Operands []int16
}

// HwDisplayOp is the operation for a display update.
type HwDisplayOp int

const (
	HwDisplayOpClearBuffer = HwDisplayOp(iota)
	HwDisplayOpSetPixel
	HwDisplayOpDisplay
)
