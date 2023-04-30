package lerp

type Remapable interface {
	Lerpable
}

type Remapper[TIn, TOut Remapable, F Float] interface {
	Remap(value TIn) TOut
	MCoeff() F
	InputMinimum() TIn
	InputMaximum() TIn
	OutputMinimum() TOut
	OutputMaximum() TOut
}

type Remapper32[TIn, TOut Remapable] Remapper[TIn, TOut, float32]

type Remapper64[TIn, TOut Remapable] Remapper[TIn, TOut, float64]
