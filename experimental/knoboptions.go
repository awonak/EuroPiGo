package experimental

import "fmt"

type KnobOption func(e *knobBankEntry) error

func InitialPercentageValue(v float32) KnobOption {
	return func(e *knobBankEntry) error {
		if v < 0 || v > 1 {
			return fmt.Errorf("initial percentage value of %f is outside the range [0..1]", v)
		}
		e.percent = v
		return nil
	}
}
