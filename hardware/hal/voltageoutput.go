package hal

import (
	"time"

	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/units"
)

type VoltageOutput interface {
	SetVoltage(v float32)
	SetCV(cv units.CV)
	SetBipolarCV(cv units.BipolarCV)
	SetVOct(voct units.VOct)
	Voltage() float32
	MinVoltage() float32
	MaxVoltage() float32
}

type VoltageOutputConfig struct {
	Period      time.Duration
	Calibration envelope.Map[float32, uint16]
}
