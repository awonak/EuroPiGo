package knobbank

import (
	"math"

	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/lerp"
)

type knobBankEntry struct {
	name       string
	enabled    bool
	locked     bool
	value      float32
	percent    float32
	vlerp      lerp.Lerper32[float32]
	minVoltage float32
	maxVoltage float32
	scale      float32
}

func (e *knobBankEntry) lock(knob hal.KnobInput, lastValue float32) float32 {
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
	return clamp.Clamp(e.percent*e.scale, 0, 1)
}

func (e *knobBankEntry) Value() float32 {
	return clamp.Clamp(e.value*e.scale, e.minVoltage, e.maxVoltage)
}

func (e *knobBankEntry) update(percent, value, lastValue float32) float32 {
	if !e.enabled || e.locked {
		return lastValue
	}

	if math.Abs(float64(value-lastValue)) < 0.05 {
		return lastValue
	}

	e.percent = percent
	e.value = value
	return value
}
