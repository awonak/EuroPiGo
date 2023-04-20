package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/randomskips/module"
	"github.com/heucuva/europi/units"
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

func (m *Settings) Button1(e *europi.EuroPi, p machine.Pin) {
	m.km.Next()
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Paint(e, deltaTime)
}
