package knobbank

import (
	"github.com/heucuva/europi/input"
	"github.com/heucuva/europi/math"
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type KnobBank struct {
	knob      input.AnalogReader
	current   int
	lastValue float32
	bank      []knobBankEntry
}

func NewKnobBank(knob input.AnalogReader, opts ...KnobBankOption) (*KnobBank, error) {
	kb := &KnobBank{
		knob:      knob,
		lastValue: knob.ReadVoltage(),
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

func (kb *KnobBank) CurrentIndex() int {
	return kb.current
}

func (kb *KnobBank) Current() input.AnalogReader {
	return kb
}

func (kb *KnobBank) Samples(samples uint16) {
	kb.knob.Samples(samples)
}

func (kb *KnobBank) ReadVoltage() float32 {
	value := kb.knob.ReadVoltage()
	if len(kb.bank) == 0 {
		return value
	}

	cur := &kb.bank[kb.current]
	percent := kb.knob.Percent()
	kb.lastValue = cur.update(percent, value, kb.lastValue)
	return cur.Value()
}

func (kb *KnobBank) ReadCV() units.CV {
	return units.CV(math.Clamp(kb.Percent(), 0.0, 1.0))
}

func (kb *KnobBank) ReadVOct() units.VOct {
	return units.VOct(kb.ReadVoltage())
}

func (kb *KnobBank) Percent() float32 {
	percent := kb.knob.Percent()
	if len(kb.bank) == 0 {
		return percent
	}

	cur := &kb.bank[kb.current]
	value := kb.knob.ReadVoltage()
	kb.lastValue = cur.update(percent, value, kb.lastValue)
	return cur.Percent()
}

func (kb *KnobBank) Range(steps uint16) uint16 {
	return kb.knob.Range(steps)
}

func (kb *KnobBank) Choice(numItems int) int {
	if len(kb.bank) == 0 {
		return int(kb.Range(uint16(numItems)))
	}

	cur := &kb.bank[kb.current]
	value := kb.knob.ReadVoltage()
	percent := kb.knob.Percent()
	kb.lastValue = cur.update(percent, value, kb.lastValue)
	idx := europim.Lerp(cur.Percent(), 0, 2*numItems+1) / 2
	return europim.Clamp(idx, 0, numItems-1)
}

func (kb *KnobBank) Next() {
	if len(kb.bank) == 0 {
		kb.current = 0
		return
	}

	cur := &kb.bank[kb.current]
	cur.lock(kb.knob, kb.lastValue)

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

func (e *knobBankEntry) lock(knob input.AnalogReader, lastValue float32) float32 {
	if e.locked {
		return lastValue
	}

	e.locked = true
	value := knob.ReadVoltage()
	percent := knob.Percent()
	return e.update(percent, value, lastValue)
}

func (e *knobBankEntry) unlock() {
	if !e.enabled {
		return
	}

	e.locked = false
}

func (e *knobBankEntry) Percent() float32 {
	return europim.Lerp[float32](e.percent*e.scale, 0, 1)
}

func (e *knobBankEntry) Value() float32 {
	return europim.Clamp(e.value*e.scale, e.minVoltage, e.maxVoltage)
}

func (e *knobBankEntry) update(percent, value, lastValue float32) float32 {
	if !e.enabled || e.locked {
		return lastValue
	}

	if europim.Abs(value-lastValue) < 0.05 {
		return lastValue
	}

	e.percent = percent
	e.value = value
	return value
}
