package envelope

import (
	"sort"

	"github.com/awonak/EuroPiGo/lerp"
)

type envMap32[TIn, TOut lerp.Lerpable] struct {
	rem    []lerp.Remapper32[TIn, TOut]
	outMax TOut
}

func NewMap32[TIn, TOut lerp.Lerpable](points []MapEntry[TIn, TOut]) Map[TIn, TOut] {
	if len(points) == 0 {
		panic("must have at least 1 point")
	}

	p := make(MapEntryList[TIn, TOut], len(points))
	// make a copy just in case we're dealing with another goroutine's data
	copy(p, points)
	// ensure it's sorted
	sort.Sort(p)

	var outMax TOut
	var rem []lerp.Remapper32[TIn, TOut]
	if len(p) == 1 {
		cur := p[0]
		outMax = cur.Output
		rem = append(rem, lerp.NewRemapPoint[TIn, TOut, float32](cur.Input, cur.Output))
	} else {
		for pos := 0; pos < len(p)-1; pos++ {
			cur, next := p[pos], p[pos+1]
			outMax = next.Output
			rem = append(rem, lerp.NewRemap32(cur.Input, next.Input, cur.Output, next.Output))
		}
	}
	return &envMap32[TIn, TOut]{
		rem:    rem,
		outMax: outMax,
	}
}

func (m *envMap32[TIn, TOut]) Remap(value TIn) TOut {
	for _, r := range m.rem {
		if value < r.InputMinimum() {
			return r.OutputMinimum()
		} else if value < r.InputMaximum() {
			return r.Remap(value)
		}
	}

	return m.outMax
}

func (m *envMap32[TIn, TOut]) InputMinimum() TIn {
	// we're guaranteed to have 1 point
	return m.rem[0].InputMaximum()
}

func (m *envMap32[TIn, TOut]) InputMaximum() TIn {
	// we're guaranteed to have 1 point
	return m.rem[len(m.rem)-1].InputMaximum()
}

func (m *envMap32[TIn, TOut]) OutputMinimum() TOut {
	// we're guaranteed to have 1 point
	return m.rem[0].OutputMinimum()
}

func (m *envMap32[TIn, TOut]) OutputMaximum() TOut {
	return m.outMax
}
