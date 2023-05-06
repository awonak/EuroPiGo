package lerp

type remap64[TIn, TOut Lerpable] struct {
	inLerp  Lerper64[TIn]
	outLerp Lerper64[TOut]
}

func NewRemap64[TIn, TOut Lerpable](inMin, inMax TIn, outMin, outMax TOut) Remapper64[TIn, TOut] {
	return remap64[TIn, TOut]{
		inLerp:  NewLerp64(inMin, inMax),
		outLerp: NewLerp64(outMin, outMax),
	}
}

func (r remap64[TIn, TOut]) Remap(value TIn) TOut {
	t := r.inLerp.InverseLerp(value)
	return r.outLerp.Lerp(t)
}

func (r remap64[TIn, TOut]) Unmap(value TOut) TIn {
	t := r.outLerp.InverseLerp(value)
	return r.inLerp.Lerp(t)
}

func (r remap64[TIn, TOut]) InputMinimum() TIn {
	return r.inLerp.OutputMinimum()
}

func (r remap64[TIn, TOut]) InputMaximum() TIn {
	return r.inLerp.OutputMaximum()
}

func (r remap64[TIn, TOut]) OutputMinimum() TOut {
	return r.outLerp.OutputMinimum()
}

func (r remap64[TIn, TOut]) OutputMaximum() TOut {
	return r.outLerp.OutputMaximum()
}
