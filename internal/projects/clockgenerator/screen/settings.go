package screen

import (
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	km    *knobmenu.KnobMenu
	Clock *module.ClockGenerator
}

func (m *Settings) bpmString() string {
	return module.BPMString(m.Clock.BPM())
}

func (m *Settings) bpmValue() units.CV {
	return module.BPMToCV(m.Clock.BPM())
}

func (m *Settings) setBPMValue(value units.CV) {
	m.Clock.SetBPM(module.CVToBPM(value))
}

func (m *Settings) gateDurationString() string {
	return module.GateDurationString(m.Clock.GateDuration())
}

func (m *Settings) gateDurationValue() units.CV {
	return module.GateDurationToCV(m.Clock.GateDuration())
}

func (m *Settings) setGateDurationValue(value units.CV) {
	m.Clock.SetGateDuration(module.CVToGateDuration(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("bpm", "BPM", m.bpmString, m.bpmValue, m.setBPMValue),
		knobmenu.WithItem("gateDuration", "Gate", m.gateDurationString, m.gateDurationValue, m.setGateDurationValue),
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
