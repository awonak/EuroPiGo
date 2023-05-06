package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/lerp"
)

const (
	// DefaultCalibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	DefaultCalibratedMinAI = 300
	DefaultCalibratedMaxAI = 44009

	DefaultSamples = 1000

	MaxInputVoltage = 10.0
	MinInputVoltage = 0.0
)

var (
	DefaultAICalibration = lerp.NewRemap32[uint16, float32](DefaultCalibratedMinAI, DefaultCalibratedMaxAI, MinInputVoltage, MaxInputVoltage)

	aiInitialConfig = hal.AnalogInputConfig{
		Samples:     DefaultSamples,
		Calibration: DefaultAICalibration,
	}
)
