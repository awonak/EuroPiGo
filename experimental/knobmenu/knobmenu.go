package knobmenu

import (
	"fmt"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobbank"
	"github.com/heucuva/europi/input"
	europim "github.com/heucuva/europi/math"
)

type KnobMenu struct {
	kb             *knobbank.KnobBank
	items          []item
	selectedRune   rune
	unselectedRune rune
	x              int16
	y              int16
	yadvance       int16
}

func NewKnobMenu(knob input.AnalogReader, opts ...KnobMenuOption) (*KnobMenu, error) {
	km := &KnobMenu{
		selectedRune:   '*',
		unselectedRune: ' ',
		x:              0,
		y:              11,
		yadvance:       12,
	}

	kbopts := []knobbank.KnobBankOption{
		knobbank.WithDisabledKnob(),
	}

	for _, opt := range opts {
		kbo, err := opt(km)
		if err != nil {
			return nil, err
		}

		kbopts = append(kbopts, kbo...)
	}

	kb, err := knobbank.NewKnobBank(knob, kbopts...)
	if err != nil {
		return nil, err
	}

	km.kb = kb

	return km, nil
}

func (m *KnobMenu) Next() {
	m.kb.Next()
}

func (m *KnobMenu) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.updateMenu(e)

	disp := e.Display

	y := m.y
	selectedIdx := m.kb.CurrentIndex() - 1
	minI := europim.Clamp(selectedIdx-1, 0, len(m.items)-1)
	maxI := europim.Clamp(minI+1, 0, len(m.items)-1)
	for i := minI; i <= maxI && i < len(m.items); i++ {
		it := &m.items[i]

		selRune := m.unselectedRune
		if i == selectedIdx {
			selRune = m.selectedRune
		}

		disp.WriteLine(fmt.Sprintf("%c%s:%s", selRune, it.label, it.stringFn()), m.x, y)
		y += m.yadvance
	}
}

func (m *KnobMenu) updateMenu(e *europi.EuroPi) {
	cur := m.kb.CurrentName()
	for _, it := range m.items {
		if it.name == cur {
			it.updateFn(m.kb.ReadCV())
			return
		}
	}
}
