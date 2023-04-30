package hal

import (
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/units"
)

type AnalogInput interface {
	Configure(config AnalogInputConfig) error
	Percent() float32
	ReadRawVoltage() uint16
	ReadVoltage() float32
	ReadCV() units.CV
	ReadBipolarCV() units.BipolarCV
	ReadVOct() units.VOct
	MinVoltage() float32
	MaxVoltage() float32
}

type AnalogInputConfig struct {
	Samples     int
	Calibration envelope.Map[uint16, float32]
}
