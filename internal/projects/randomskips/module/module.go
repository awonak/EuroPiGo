package module

import (
	"math/rand"
	"time"

	"github.com/awonak/EuroPiGo/units"
)

type RandomSkips struct {
	gate   func(high bool)
	chance float32

	active    bool
	lastInput bool
	cv        float32
	ac        float32 // attenuated chance (cv * chance)
}

func (m *RandomSkips) Init(config Config) error {
	fnGate := config.Gate
	if fnGate == nil {
		fnGate = noopGate
	}
	m.gate = fnGate
	m.chance = config.Chance

	m.SetCV(1)
	return nil
}

func noopGate(high bool) {
}

func (m *RandomSkips) Gate(value bool) {
	prev := m.active
	lastInput := m.lastInput
	next := prev
	m.lastInput = value

	if value != lastInput && rand.Float32() < m.ac {
		next = !prev
	}

	if prev != next {
		m.active = next
		m.gate(next)
	}
}

func (m *RandomSkips) SetChance(chance float32) {
	m.chance = chance
}

func (m *RandomSkips) Chance() float32 {
	return m.chance
}

func (m *RandomSkips) SetCV(cv units.CV) {
	m.cv = cv.ToFloat32()
	m.ac = m.chance * m.cv
}

func (m *RandomSkips) Tick(deltaTime time.Duration) {
}
