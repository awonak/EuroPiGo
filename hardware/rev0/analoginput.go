package rev0

import (
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

const (
	// DefaultCalibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	DefaultCalibratedMinAI = 300
	DefaultCalibratedMaxAI = 44009

	DefaultSamples = 1000

	MaxInputVoltage = 3.3
	MinInputVoltage = 0.0
)

var (
	aiInitialConfig = hal.AnalogInputConfig{
		Samples: DefaultSamples,
		Calibration: envelope.NewMap32([]envelope.MapEntry[uint16, float32]{
			{
				Input:  DefaultCalibratedMinAI,
				Output: MinInputVoltage,
			},
			{
				Input:  DefaultCalibratedMaxAI,
				Output: MaxInputVoltage,
			},
		}),
	}
)
