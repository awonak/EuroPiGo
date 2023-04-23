package quantizer

import (
	"github.com/heucuva/europi/lerp"
)

type Trunc[T any] struct{}

func (Trunc[T]) QuantizeToIndex(in float32, length int) int {
	if length == 0 {
		return -1
	}

	return lerp.NewLerp32(0, length-1).ClampedLerp(in)
}

func (q Trunc[T]) QuantizeToValue(in float32, list []T) T {
	idx := q.QuantizeToIndex(in, len(list))
	if idx == -1 {
		var empty T
		return empty
	}

	return list[idx]
}
