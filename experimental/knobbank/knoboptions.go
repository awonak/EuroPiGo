package knobbank

import (
	"fmt"

	"github.com/heucuva/europi/lerp"
)

type KnobOption func(e *knobBankEntry) error

func InitialPercentageValue(v float32) KnobOption {
	return func(e *knobBankEntry) error {
		if v < 0 || v > 1 {
			return fmt.Errorf("initial percentage value of %f is outside the range [0..1]", v)
		}

		e.percent = v
		e.vlerp = lerp.NewLerp32[float32](defaultMinInputVoltage, defaultMaxInputVoltage)
		e.value = e.vlerp.ClampedLerp(v)
		return nil
	}
}

func MinInputVoltage(v float32) KnobOption {
	return func(e *knobBankEntry) error {
		e.minVoltage = v
		e.vlerp = lerp.NewLerp32(e.minVoltage, e.maxVoltage)
		e.value = e.vlerp.ClampedLerp(e.percent)
		return nil
	}
}

func MaxInputVoltage(v float32) KnobOption {
	return func(e *knobBankEntry) error {
		e.maxVoltage = v
		e.vlerp = lerp.NewLerp32(e.minVoltage, e.maxVoltage)
		e.value = e.vlerp.ClampedLerp(e.percent)
		return nil
	}
}
