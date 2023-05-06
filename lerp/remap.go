package lerp

type Remapper[TIn, TOut Lerpable, TFloat Float] interface {
	Remap(in TIn) TOut
	Unmap(out TOut) TIn
	InputMinimum() TIn
	InputMaximum() TIn
	OutputMinimum() TOut
	OutputMaximum() TOut
}

type Remapper32[TIn, TOut Lerpable] Remapper[TIn, TOut, float32]
type Remapper64[TIn, TOut Lerpable] Remapper[TIn, TOut, float64]
