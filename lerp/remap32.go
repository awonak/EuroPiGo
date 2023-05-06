package lerp

type remap32[TIn, TOut Lerpable] struct {
	inLerp  Lerper32[TIn]
	outLerp Lerper32[TOut]
}

func NewRemap32[TIn, TOut Lerpable](inMin, inMax TIn, outMin, outMax TOut) Remapper32[TIn, TOut] {
	return remap32[TIn, TOut]{
		inLerp:  NewLerp32(inMin, inMax),
		outLerp: NewLerp32(outMin, outMax),
	}
}

func (r remap32[TIn, TOut]) Remap(value TIn) TOut {
	t := r.inLerp.InverseLerp(value)
	return r.outLerp.Lerp(t)
}

func (r remap32[TIn, TOut]) Unmap(value TOut) TIn {
	t := r.outLerp.InverseLerp(value)
	return r.inLerp.Lerp(t)
}

func (r remap32[TIn, TOut]) InputMinimum() TIn {
	return r.inLerp.OutputMinimum()
}

func (r remap32[TIn, TOut]) InputMaximum() TIn {
	return r.inLerp.OutputMaximum()
}

func (r remap32[TIn, TOut]) OutputMinimum() TOut {
	return r.outLerp.OutputMinimum()
}

func (r remap32[TIn, TOut]) OutputMaximum() TOut {
	return r.outLerp.OutputMaximum()
}
