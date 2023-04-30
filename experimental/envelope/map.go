package envelope

import "github.com/awonak/EuroPiGo/lerp"

type Map[TIn, TOut lerp.Lerpable] interface {
	Remap(value TIn) TOut
	InputMinimum() TIn
	InputMaximum() TIn
	OutputMinimum() TOut
	OutputMaximum() TOut
}
