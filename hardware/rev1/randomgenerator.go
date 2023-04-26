package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type randomGenerator struct {
	rnd RNDProvider
}

var (
	// static check
	_ hal.RandomGenerator = &randomGenerator{}
	// silence linter
	_ = newRandomGenerator
)

func newRandomGenerator(rnd RNDProvider) hal.RandomGenerator {
	return &randomGenerator{
		rnd: rnd,
	}
}

type RNDProvider interface {
	Configure(config hal.RandomGeneratorConfig) error
}

// Configure updates the device with various configuration parameters
func (r *randomGenerator) Configure(config hal.RandomGeneratorConfig) error {
	if r.rnd != nil {
		if err := r.rnd.Configure(config); err != nil {
			return err
		}
	}

	return nil
}
