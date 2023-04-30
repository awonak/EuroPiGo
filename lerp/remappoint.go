package lerp

// This is the world's worst Lerp remapper. Regardless of the input value, it always returns the output value.

type remapPoint[TIn, TOut Remapable, TFloat Float] struct {
	in  TIn
	out TOut
}

func NewRemapPoint[TIn, TOut Remapable, TFloat Float](in TIn, out TOut) Remapper[TIn, TOut, TFloat] {
	return remapPoint[TIn, TOut, TFloat]{
		in:  in,
		out: out,
	}
}

func NewRemapPoint32[TIn, TOut Remapable](in TIn, out TOut) Remapper[TIn, TOut, float32] {
	return NewRemapPoint[TIn, TOut, float32](in, out)
}

func NewRemapPoint64[TIn, TOut Remapable](in TIn, out TOut) Remapper[TIn, TOut, float64] {
	return NewRemapPoint[TIn, TOut, float64](in, out)
}

func (r remapPoint[TIn, TOut, TFloat]) Remap(value TIn) TOut {
	// `value` isn't used here - just return `out`
	return r.out
}

func (r remapPoint[TIn, TOut, TFloat]) MCoeff() TFloat {
	return 0.0
}

func (r remapPoint[TIn, TOut, TFloat]) InputMinimum() TIn {
	return r.in
}

func (r remapPoint[TIn, TOut, TFloat]) InputMaximum() TIn {
	return r.in
}

func (r remapPoint[TIn, TOut, TFloat]) OutputMinimum() TOut {
	return r.out
}

func (r remapPoint[TIn, TOut, TFloat]) OutputMaximum() TOut {
	return r.out
}
