package common

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type RandomGenerator struct {
	rnd RNDProvider
}

var (
	// static check
	_ hal.RandomGenerator = (*RandomGenerator)(nil)
	// silence linter
	_ = NewRandomGenerator
)

func NewRandomGenerator(rnd RNDProvider) *RandomGenerator {
	return &RandomGenerator{
		rnd: rnd,
	}
}

type RNDProvider interface {
	Configure(config hal.RandomGeneratorConfig) error
}

// Configure updates the device with various configuration parameters
func (r *RandomGenerator) Configure(config hal.RandomGeneratorConfig) error {
	if r.rnd != nil {
		if err := r.rnd.Configure(config); err != nil {
			return err
		}
	}

	return nil
}
