package envelope

import "github.com/awonak/EuroPiGo/lerp"

type Map[TIn, TOut lerp.Lerpable] interface {
	Remap(value TIn) TOut
	Unmap(value TOut) TIn
	InputMinimum() TIn
	InputMaximum() TIn
	OutputMinimum() TOut
	OutputMaximum() TOut
}

type remapList[TIn, TOut lerp.Lerpable, TFloat lerp.Float] struct {
	lerp.Remapper[TIn, TOut, TFloat]
	nextOut *remapList[TIn, TOut, TFloat]
}
