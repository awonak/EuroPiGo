package hal

import (
	"time"

	"github.com/awonak/EuroPiGo/lerp"
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
	Monopolar   bool
	Calibration lerp.Remapper32[float32, uint16]
}
