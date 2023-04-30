package envelope

import "github.com/awonak/EuroPiGo/lerp"

type MapEntry[TIn, TOut lerp.Lerpable] struct {
	Input  TIn
	Output TOut
}

type MapEntryList[TIn, TOut lerp.Lerpable] []MapEntry[TIn, TOut]

func (m MapEntryList[TIn, TOut]) Len() int {
	return len(m)
}

func (m MapEntryList[TIn, TOut]) Less(i, j int) bool {
	return m[i].Input < m[j].Input
}

func (m MapEntryList[TIn, TOut]) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
