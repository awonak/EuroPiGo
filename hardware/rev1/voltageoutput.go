package rev1

import (
	"time"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/lerp"
)

const (
	// Manually calibrated to best match expected voltages. Additional info:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	CalibratedOffset = 0
	// The default pwmGroup Top of MaxUint16 caused noisy output. Dropping this down to a 8bit value resulted in much smoother cv output.
	CalibratedTop = 0xff - CalibratedOffset

	MaxOutputVoltage = 10.0
	MinOutputVoltage = 0.0

	// We need a rather high frequency to achieve a stable cv ouput, which means we need a rather low duty cycle period.
	// Set a period of 500ns.
	DefaultPWMPeriod time.Duration = time.Nanosecond * 500
)

var (
	DefaultVoltageOutputCalibration = lerp.NewRemap32[float32, uint16](MinOutputVoltage, MaxOutputVoltage, CalibratedOffset, CalibratedTop)

	cvInitialConfig = hal.VoltageOutputConfig{
		Period:      DefaultPWMPeriod,
		Monopolar:   true,
		Calibration: DefaultVoltageOutputCalibration,
	}
)
