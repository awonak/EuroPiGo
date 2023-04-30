package lerp

type remap32[TIn, TOut Remapable] struct {
	inMin  TIn
	inMax  TIn
	outMin TOut
	r      float32
}

func NewRemap32[TIn, TOut Remapable](inMin, inMax TIn, outMin, outMax TOut) Remapper32[TIn, TOut] {
	var r float32
	// if rIn is 0, then we don't need to test further, we're always min (max) value
	if rIn := inMax - inMin; rIn != 0 {
		if rOut := outMax - outMin; rOut != 0 {
			r = float32(rOut) / float32(rIn)
		}
	}
	return remap32[TIn, TOut]{
		inMin:  inMin,
		inMax:  inMax,
		outMin: outMin,
		r:      r,
	}
}

func (r remap32[TIn, TOut]) Remap(value TIn) TOut {
	if r.r == 0.0 {
		return r.outMin
	}

	return r.outMin + TOut(r.r*float32(value-r.inMin))
}

func (r remap32[TIn, TOut]) MCoeff() float32 {
	return r.r
}

func (r remap32[TIn, TOut]) InputMinimum() TIn {
	return r.inMin
}

func (r remap32[TIn, TOut]) InputMaximum() TIn {
	return r.inMax
}

func (r remap32[TIn, TOut]) OutputMinimum() TOut {
	return r.outMin
}

func (r remap32[TIn, TOut]) OutputMaximum() TOut {
	return r.Remap(r.inMax)
}
