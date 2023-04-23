package hal

import (
	"time"

	"github.com/heucuva/europi/units"
)

type VoltageOutput interface {
	SetVoltage(v float32)
	SetCV(cv units.CV)
	SetVOct(voct units.VOct)
	Voltage() float32
	MinVoltage() float32
	MaxVoltage() float32
}

type VoltageOutputConfig struct {
	Period time.Duration
	Offset uint16
	Top    uint16
}
