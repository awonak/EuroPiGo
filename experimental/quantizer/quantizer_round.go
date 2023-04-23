package quantizer

import (
	"github.com/awonak/EuroPiGo/lerp"
)

type Round[T any] struct{}

func (Round[T]) QuantizeToIndex(in float32, length int) int {
	if length == 0 {
		return -1
	}

	return lerp.NewLerp32(0, length-1).ClampedLerpRound(in)
}

func (q Round[T]) QuantizeToValue(in float32, list []T) T {
	idx := q.QuantizeToIndex(in, len(list))
	if idx == -1 {
		var empty T
		return empty
	}

	return list[idx]
}
