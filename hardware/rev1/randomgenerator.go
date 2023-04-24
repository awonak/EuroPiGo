package rev1

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
)

type randomGenerator struct {
	rnd rndProvider
}

var (
	// static check
	_ hal.RandomGenerator = &randomGenerator{}
	// silence linter
	_ = newRandomGenerator
)

func newRandomGenerator(rnd rndProvider) hal.RandomGenerator {
	return &randomGenerator{
		rnd: rnd,
	}
}

type rndProvider interface {
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
