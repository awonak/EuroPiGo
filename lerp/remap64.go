package lerp

type remap64[TIn, TOut Remapable] struct {
	inMin  TIn
	inMax  TIn
	outMin TOut
	outMax TOut
	r      float64
}

func NewRemap64[TIn, TOut Remapable](inMin, inMax TIn, outMin, outMax TOut) Remapper64[TIn, TOut] {
	var r float64
	// if rIn is 0, then we don't need to test further, we're always min (max) value
	if rIn := float64(inMax) - float64(inMin); rIn != 0 {
		if rOut := float64(outMax) - float64(outMin); rOut != 0 {
			r = rOut / rIn
		}
	}
	return remap64[TIn, TOut]{
		inMin:  inMin,
		inMax:  inMax,
		outMin: outMin,
		outMax: outMax,
		r:      r,
	}
}

func (r remap64[TIn, TOut]) Remap(value TIn) TOut {
	switch {
	case r.r == 0.0:
		return r.outMin
	case value == r.inMin:
		return r.outMin
	case value == r.inMax:
		return r.outMax
	default:
		return r.outMin + TOut(r.r*float64(value-r.inMin))
	}
}

func (r remap64[TIn, TOut]) Unmap(value TOut) TIn {
	if r.r == 0.0 {
		return r.inMax
	}

	rOut := float64(value) - float64(r.outMin)
	return r.inMin + TIn(rOut/r.r)
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
