package quantizer

import (
	"github.com/awonak/EuroPiGo/lerp"
)

// Round is a rounding-style quantizer
type Round[T any] struct{}

// QuantizeToIndex takes a normalized input value and a length value, then provides
// a return value between 0 and (length - 1), inclusive.
//
// A return value of -1 means that the length parameter was 0 (a value that cannot be quantized over successfully).
func (Round[T]) QuantizeToIndex(in float32, length int) int {
	if length == 0 {
		return -1
	}

	return lerp.NewLerp32(0, length-1).ClampedLerpRound(in)
}

// QuantizeToValue takes a normalized input value and a list of values, then provides
// a return value chosen from the provided list of values.
//
// A return value of the zeroish equivalent of the value means that the list parameter
// was empty (this situation does not lend well to quantization).
func (q Round[T]) QuantizeToValue(in float32, list []T) T {
	idx := q.QuantizeToIndex(in, len(list))
	if idx == -1 {
		var empty T
		return empty
	}

	return list[idx]
}
