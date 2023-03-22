package experimental

import (
	"fmt"

	"github.com/heucuva/europi/input"
	"github.com/heucuva/europi/internal/math"
)

type KnobOption func(e *knobBankEntry) error

func InitialPercentageValue(v float32) KnobOption {
	return func(e *knobBankEntry) error {
		if v < 0 || v > 1 {
			return fmt.Errorf("initial percentage value of %f is outside the range [0..1]", v)
		}

		e.percent = v
		e.value = math.Lerp[float32](v, input.MinVoltage, input.MaxVoltage)
		return nil
	}
}
