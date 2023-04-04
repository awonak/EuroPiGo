package knobbank

import (
	"fmt"

	"github.com/heucuva/europi/input"
)

type KnobBankOption func(kb *KnobBank) error

func WithDisabledKnob() KnobBankOption {
	return func(kb *KnobBank) error {
		kb.bank = append(kb.bank, knobBankEntry{})
		return nil
	}
}

func WithLockedKnob(name string, opts ...KnobOption) KnobBankOption {
	return func(kb *KnobBank) error {
		e := knobBankEntry{
			name:       name,
			enabled:    true,
			locked:     true,
			minVoltage: input.MinVoltage,
			maxVoltage: input.MaxVoltage,
			scale:      1,
		}

		for _, opt := range opts {
			if err := opt(&e); err != nil {
				return fmt.Errorf("%s knob configuration error: %w", name, err)
			}
		}

		kb.bank = append(kb.bank, e)
		return nil
	}
}
