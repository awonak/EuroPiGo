package knobmenu

import (
	"fmt"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	"github.com/awonak/EuroPiGo/experimental/knobbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"tinygo.org/x/tinyfont/proggy"
)

var (
	DefaultFont = &proggy.TinySZ8pt7b
)

type KnobMenu struct {
	kb             *knobbank.KnobBank
	items          []item
	selectedRune   rune
	unselectedRune rune
	x              int16
	y              int16
	yadvance       int16
	writer         fontwriter.Writer
}

func NewKnobMenu(knob hal.KnobInput, opts ...KnobMenuOption) (*KnobMenu, error) {
	km := &KnobMenu{
		selectedRune:   '*',
		unselectedRune: ' ',
		x:              0,
		y:              11,
		yadvance:       12,
		writer: fontwriter.Writer{
			Display: nil,
			Font:    DefaultFont,
		},
	}

	km.yadvance = int16(km.writer.Font.GetYAdvance())
	km.y = km.yadvance

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

func (m *KnobMenu) Paint(e europi.Hardware, deltaTime time.Duration) {
	m.updateMenu(e)

	m.writer.Display = europi.Display(e)
	if m.writer.Display == nil {
		return
	}

	y := m.y
	selectedIdx := m.kb.CurrentIndex() - 1
	minI := clamp.Clamp(selectedIdx-1, 0, len(m.items)-1)
	maxI := clamp.Clamp(minI+1, 0, len(m.items)-1)
	for i := minI; i <= maxI && i < len(m.items); i++ {
		it := &m.items[i]

		selRune := m.unselectedRune
		if i == selectedIdx {
			selRune = m.selectedRune
		}

		m.writer.WriteLine(fmt.Sprintf("%c%s:%s", selRune, it.label, it.stringFn()), m.x, y, draw.White)
		y += m.yadvance
	}
}

func (m *KnobMenu) updateMenu(e europi.Hardware) {
	cur := m.kb.CurrentName()
	for _, it := range m.items {
		if it.name == cur {
			it.updateFn(m.kb.ReadCV())
			return
		}
	}
}
