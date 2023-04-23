package rev1

import "github.com/heucuva/europi/internal/hardware/hal"

type HwMessageDigitalValue struct {
	Value bool
}

type HwMessageADCValue struct {
	Value uint16
}

type HwMessageInterrupt struct {
	Change hal.ChangeFlags
}

type HwMessagePwmValue struct {
	Value uint16
}

type HwMessageDisplay struct {
	Op       HwDisplayOp
	Operands []int16
}

type HwDisplayOp int

const (
	HwDisplayOpClearBuffer = HwDisplayOp(iota)
	HwDisplayOpSetPixel
	HwDisplayOpDisplay
)
