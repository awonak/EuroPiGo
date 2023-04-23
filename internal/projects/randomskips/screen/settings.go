package screen

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/knobmenu"
	"github.com/awonak/EuroPiGo/internal/projects/randomskips/module"
	"github.com/awonak/EuroPiGo/units"
)

type Settings struct {
	km          *knobmenu.KnobMenu
	RandomSkips *module.RandomSkips
}

func (m *Settings) chanceString() string {
	return module.ChanceString(m.RandomSkips.Chance())
}

func (m *Settings) chanceValue() units.CV {
	return module.ChanceToCV(m.RandomSkips.Chance())
}

func (m *Settings) setChanceValue(value units.CV) {
	m.RandomSkips.SetChance(module.CVToChance(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("chance", "Chance", m.chanceString, m.chanceValue, m.setChanceValue),
	)
	if err != nil {
		panic(err)
	}

	m.km = km
}

func (m *Settings) Button1Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Settings) Button1(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Next()
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Paint(e, deltaTime)
}
