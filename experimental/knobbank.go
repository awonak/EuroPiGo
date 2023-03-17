package experimental

import (
	"github.com/heucuva/europi/input"
	"github.com/heucuva/europi/internal/math"
)

type KnobBank struct {
	knob    input.AnalogReader
	current int
	bank    []knobBankEntry
}

func NewKnobBank(knob input.AnalogReader, opts ...KnobBankOption) (*KnobBank, error) {
	kb := &KnobBank{
		knob: knob,
	}

	for _, opt := range opts {
		if err := opt(kb); err != nil {
			return nil, err
		}
	}

	return kb, nil
}

func (kb *KnobBank) CurrentName() string {
	if len(kb.bank) == 0 {
		return ""
	}
	return kb.bank[kb.current].name
}

func (kb *KnobBank) Current() input.AnalogReader {
	return kb
}

func (kb *KnobBank) Samples(samples uint16) {
	kb.knob.Samples(samples)
}

func (kb *KnobBank) ReadVoltage() float32 {
	if len(kb.bank) == 0 {
		return kb.knob.ReadVoltage()
	}

	cur := &kb.bank[kb.current]
	cur.update(kb.knob)
	return cur.value
}

func (kb *KnobBank) Percent() float32 {
	if len(kb.bank) == 0 {
		return kb.knob.Percent()
	}

	cur := &kb.bank[kb.current]
	cur.update(kb.knob)
	return cur.percent
}

func (kb *KnobBank) Range(steps uint16) uint16 {
	return kb.knob.Range(steps)
}

func (kb *KnobBank) Choice(numItems int) int {
	if len(kb.bank) == 0 {
		return int(kb.Range(uint16(numItems)))
	}

	cur := &kb.bank[kb.current]
	cur.update(kb.knob)
	return math.Lerp(cur.percent, 0, numItems-1)
}

func (kb *KnobBank) Next() {
	if len(kb.bank) == 0 {
		kb.current = 0
		return
	}

	kb.bank[kb.current].lock(kb.knob)
	kb.current++
	if kb.current >= len(kb.bank) {
		kb.current = 0
	}
	kb.bank[kb.current].unlock()
}

type knobBankEntry struct {
	name       string
	enabled    bool
	locked     bool
	value      float32
	percent    float32
	minVoltage float32
	maxVoltage float32
	scale      float32
}

func (e *knobBankEntry) lock(knob input.AnalogReader) {
	if e.locked {
		return
	}

	e.update(knob)
	e.locked = true
}

func (e *knobBankEntry) unlock() {
	if !e.enabled {
		return
	}

	e.locked = false
}

func (e *knobBankEntry) update(knob input.AnalogReader) {
	if !e.enabled || e.locked {
		return
	}

	e.percent = math.Lerp[float32](knob.Percent()*e.scale, 0, 1)
	e.value = math.Clamp(knob.ReadVoltage()*e.scale, e.minVoltage, e.maxVoltage)
}
