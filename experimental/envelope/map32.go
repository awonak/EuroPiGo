package envelope

import (
	"sort"

	"github.com/awonak/EuroPiGo/lerp"
)

type envMap32[TIn, TOut lerp.Lerpable] struct {
	rem     []remapList[TIn, TOut, float32]
	outMax  TOut
	outRoot *remapList[TIn, TOut, float32]
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

	var rem []remapList[TIn, TOut, float32]
	if len(p) > 1 {
		for pos := 0; pos < len(p)-1; pos++ {
			cur, next := p[pos], p[pos+1]
			rem = append(rem, remapList[TIn, TOut, float32]{
				Remapper: lerp.NewRemap32(cur.Input, next.Input, cur.Output, next.Output),
			})
		}
	}
	last := &p[len(p)-1]
	rem = append(rem, remapList[TIn, TOut, float32]{
		Remapper: lerp.NewRemapPoint[TIn, TOut, float32](last.Input, last.Output),
	})

	outSort := make(MapEntryList[TOut, int], len(rem))
	for i, e := range rem {
		outSort[i].Input = e.OutputMinimum()
		outSort[i].Output = i
	}
	sort.Sort(outSort)
	rootIdx := outSort[0].Output
	outRoot := &rem[rootIdx]
	for pos := 0; pos < len(rem)-1; pos++ {
		cur, next := outSort[pos].Output, outSort[pos+1].Output
		rem[cur].nextOut = &rem[next]
	}

	return &envMap32[TIn, TOut]{
		rem:     rem,
		outMax:  last.Output,
		outRoot: outRoot,
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

func (m *envMap32[TIn, TOut]) Unmap(value TOut) TIn {
	for r := m.outRoot; r != nil; r = r.nextOut {
		outMin := r.OutputMinimum()
		outMax := r.OutputMaximum()
		if outMin < outMax {
			if value < outMin {
				return r.InputMinimum()
			} else if value < outMax {
				return r.Unmap(value)
			}
		} else {
			if value < outMax {
				return r.InputMinimum()
			} else if value < outMin {
				return r.Unmap(value)
			}
		}
	}

	return m.InputMaximum()
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
