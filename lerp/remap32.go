package lerp

type remap32[TIn, TOut Remapable] struct {
	inMin  TIn
	inMax  TIn
	outMin TOut
	outMax TOut
	r      float32
}

func NewRemap32[TIn, TOut Remapable](inMin, inMax TIn, outMin, outMax TOut) Remapper32[TIn, TOut] {
	var r float32
	// if rIn is 0, then we don't need to test further, we're always min (max) value
	if rIn := float32(inMax) - float32(inMin); rIn != 0 {
		if rOut := float32(outMax) - float32(outMin); rOut != 0 {
			r = rOut / rIn
		}
	}
	return remap32[TIn, TOut]{
		inMin:  inMin,
		inMax:  inMax,
		outMin: outMin,
		outMax: outMax,
		r:      r,
	}
}

func (r remap32[TIn, TOut]) Remap(value TIn) TOut {
	switch {
	case r.r == 0.0:
		return r.outMin
	case value == r.inMin:
		return r.outMin
	case value == r.inMax:
		return r.outMax
	default:
		return r.outMin + TOut(r.r*float32(value-r.inMin))
	}
}

func (r remap32[TIn, TOut]) Unmap(value TOut) TIn {
	if r.r == 0.0 {
		return r.inMax
	}

	rOut := float32(value) - float32(r.outMin)
	return r.inMin + TIn(rOut/r.r)
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
