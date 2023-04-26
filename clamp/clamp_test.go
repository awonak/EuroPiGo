package clamp_test

import (
	"testing"

	"github.com/awonak/EuroPiGo/clamp"
)

func TestClamp(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		t.Run("Min", func(t *testing.T) {
			min, max := 0, 10
			if expected, actual := min, clamp.Clamp(min, min, max); actual != expected {
				t.Fatalf("Clamp[%v, %v] Clamp: expected[%v] actual[%v]", min, max, expected, actual)
			}
		})

		t.Run("Max", func(t *testing.T) {
			min, max := 0, 10
			if expected, actual := max, clamp.Clamp(max, min, max); actual != expected {
				t.Fatalf("Clamp[%v, %v] Clamp: expected[%v] actual[%v]", min, max, expected, actual)
			}
		})
	})

	t.Run("OutOfRange", func(t *testing.T) {
		t.Run("BelowMin", func(t *testing.T) {
			min, max := 0, 10
			if expected, actual := min, clamp.Clamp(min-2, min, max); actual != expected {
				t.Fatalf("Clamp[%v, %v] Clamp: expected[%v] actual[%v]", min, max, expected, actual)
			}
		})

		t.Run("AboveMax", func(t *testing.T) {
			min, max := 0, 10
			if expected, actual := max, clamp.Clamp(max+2, min, max); actual != expected {
				t.Fatalf("Clamp[%v, %v] Clamp: expected[%v] actual[%v]", min, max, expected, actual)
			}
		})
	})
}
