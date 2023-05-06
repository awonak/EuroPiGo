package quantizer_test

import (
	"testing"

	"github.com/awonak/EuroPiGo/quantizer"
)

func TestRound(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		if actual := quantizer.New[int](quantizer.ModeRound); actual == nil {
			t.Fatalf("Quantizer[%v] New: expected[non-nil] actual[nil]", quantizer.ModeRound)
		}
	})

	t.Run("QuantizeToIndex", func(t *testing.T) {
		q := quantizer.New[int](quantizer.ModeRound)

		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min := float32(0)
				length := 10
				if expected, actual := 0, q.QuantizeToIndex(min, length); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToIndex(%v, %v): expected[%v] actual[%v]", quantizer.ModeRound, min, length, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				max := float32(1)
				length := 10
				if expected, actual := length-1, q.QuantizeToIndex(max, length); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToIndex(%v, %v): expected[%v] actual[%v]", quantizer.ModeRound, max, length, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			t.Run("BelowMin", func(t *testing.T) {
				belowMin := float32(-1)
				length := 10
				if expected, actual := 0, q.QuantizeToIndex(belowMin, length); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToIndex(%v, %v): expected[%v] actual[%v]", quantizer.ModeRound, belowMin, length, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				aboveMax := float32(2)
				length := 10
				if expected, actual := length-1, q.QuantizeToIndex(aboveMax, length); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToIndex(%v, %v): expected[%v] actual[%v]", quantizer.ModeRound, aboveMax, length, expected, actual)
				}
			})

			t.Run("EmptySet", func(t *testing.T) {
				emptySet := float32(0.5)
				length := 0
				if expected, actual := -1, q.QuantizeToIndex(emptySet, length); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToIndex(%v, %v): expected[%v] actual[%v]", quantizer.ModeRound, emptySet, length, expected, actual)
				}
			})
		})
	})

	t.Run("QuantizeToValue", func(t *testing.T) {
		q := quantizer.New[int](quantizer.ModeRound)

		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min := float32(0)
				list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
				if expected, actual := list[0], q.QuantizeToValue(min, list); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToValue(%v, %T): expected[%v] actual[%v]", quantizer.ModeRound, min, list, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				max := float32(1)
				list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
				if expected, actual := list[len(list)-1], q.QuantizeToValue(max, list); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToValue(%v, %T): expected[%v] actual[%v]", quantizer.ModeRound, max, list, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			t.Run("BelowMin", func(t *testing.T) {
				belowMin := float32(-1)
				list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
				if expected, actual := list[0], q.QuantizeToValue(belowMin, list); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToValue(%v, %T): expected[%v] actual[%v]", quantizer.ModeRound, belowMin, list, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				aboveMax := float32(2)
				list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
				if expected, actual := list[len(list)-1], q.QuantizeToValue(aboveMax, list); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToValue(%v, %T): expected[%v] actual[%v]", quantizer.ModeRound, aboveMax, list, expected, actual)
				}
			})

			t.Run("EmptySet", func(t *testing.T) {
				emptySet := float32(0.5)
				var list []int
				if expected, actual := 0, q.QuantizeToValue(emptySet, list); actual != expected {
					t.Fatalf("Quantizer[%v] QuantizeToValue(%v, %T): expected[%v] actual[%v]", quantizer.ModeRound, emptySet, list, expected, actual)
				}
			})
		})
	})
}
