//go:build pico
// +build pico

package pico

import (
	"machine"
	"math/rand"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

type picoRnd struct{}

func (r *picoRnd) Configure(config hal.RandomGeneratorConfig) error {
	xl, _ := machine.GetRNG()
	xh, _ := machine.GetRNG()
	x := int64(xh)<<32 | int64(xl)
	rand.Seed(x)
	return nil
}
