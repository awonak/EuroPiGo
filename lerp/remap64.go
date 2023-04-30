package lerp

type remap64[TIn, TOut Remapable] struct {
	inMin  TIn
	inMax  TIn
	outMin TOut
	r      float64
}

func NewRemap64[TIn, TOut Remapable](inMin, inMax TIn, outMin, outMax TOut) Remapper64[TIn, TOut] {
	var r float64
	// if rIn is 0, then we don't need to test further, we're always min (max) value
	if rIn := inMax - inMin; rIn != 0 {
		if rOut := outMax - outMin; rOut != 0 {
			r = float64(rOut) / float64(rIn)
		}
	}
	return remap64[TIn, TOut]{
		inMin:  inMin,
		inMax:  inMax,
		outMin: outMin,
		r:      r,
	}
}

func (r remap64[TIn, TOut]) Remap(value TIn) TOut {
	if r.r == 0.0 {
		return r.outMin
	}

	return r.outMin + TOut(r.r*float64(value-r.inMin))
}

func (r remap64[TIn, TOut]) MCoeff() float64 {
	return r.r
}

func (r remap64[TIn, TOut]) InputMinimum() TIn {
	return r.inMin
}

func (r remap64[TIn, TOut]) InputMaximum() TIn {
	return r.inMax
}

func (r remap64[TIn, TOut]) OutputMinimum() TOut {
	return r.outMin
}

func (r remap64[TIn, TOut]) OutputMaximum() TOut {
	return r.Remap(r.inMax)
}
