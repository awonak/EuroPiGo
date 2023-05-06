package lerp_test

import (
	"math"
	"testing"

	"github.com/awonak/EuroPiGo/lerp"
)

func TestRemap64(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		inMin, inMax := 0, 10
		outMin, outMax := float64(-math.Pi), float64(math.Pi)
		if actual := lerp.NewRemap64(inMin, inMax, outMin, outMax); actual == nil {
			t.Fatalf("Remap64[%v, %v, %v, %v] NewRemap64: expected[non-nil] actual[nil]", inMin, inMax, outMin, outMax)
		}
	})

	t.Run("Remap", func(t *testing.T) {
		t.Run("ZeroRange", func(t *testing.T) {
			inMin, inMax := 10, 10
			outMin, outMax := float64(-math.Pi), float64(math.Pi)
			l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
			if expected, actual := outMin, l.Remap(inMin); actual != expected {
				t.Fatalf("Remap64[%v, %v, %v, %v] Remap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
			}
		})
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := outMin, l.Remap(inMin); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Remap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := outMax, l.Remap(inMax); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Remap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// Remap() will work as a linear extrapolator when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := float64(-4.39822971502571), l.Remap(-2); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Remap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := float64(4.39822971502571), l.Remap(12); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Remap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})
		})
	})

	t.Run("Unmap", func(t *testing.T) {
		t.Run("ZeroRange", func(t *testing.T) {
			inMin, inMax := 10, 10
			outMin, outMax := float64(-math.Pi), float64(math.Pi)
			l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
			if expected, actual := inMin, l.Unmap(outMin); actual != expected {
				t.Fatalf("Remap64[%v, %v, %v, %v] Unmap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
			}
		})
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := inMin, l.Unmap(outMin); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Unmap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := inMax, l.Unmap(outMax); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Unmap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// Unmap() will work as a linear extrapolator when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := -2, l.Unmap(float64(-4.39822971502571)); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Unmap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				inMin, inMax := 0, 10
				outMin, outMax := float64(-math.Pi), float64(math.Pi)
				l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
				if expected, actual := 12, l.Unmap(float64(4.39822971502571)); actual != expected {
					t.Fatalf("Remap64[%v, %v, %v, %v] Unmap: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
				}
			})
		})
	})

	t.Run("InputMinimum", func(t *testing.T) {
		inMin, inMax := 0, 10
		outMin, outMax := float64(-math.Pi), float64(math.Pi)
		l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
		if expected, actual := inMin, l.InputMinimum(); actual != expected {
			t.Fatalf("Remap64[%v, %v, %v, %v] InputMinimum: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
		}
	})

	t.Run("InputMaximum", func(t *testing.T) {
		inMin, inMax := 0, 10
		outMin, outMax := float64(-math.Pi), float64(math.Pi)
		l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
		if expected, actual := inMax, l.InputMaximum(); actual != expected {
			t.Fatalf("Remap64[%v, %v, %v, %v] InputMaximum: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
		}
	})

	t.Run("OutputMinimum", func(t *testing.T) {
		inMin, inMax := 0, 10
		outMin, outMax := float64(-math.Pi), float64(math.Pi)
		l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
		if expected, actual := outMin, l.OutputMinimum(); actual != expected {
			t.Fatalf("Remap64[%v, %v, %v, %v] OutputMinimum: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
		}
	})

	t.Run("OutputMaximum", func(t *testing.T) {
		inMin, inMax := 0, 10
		outMin, outMax := float64(-math.Pi), float64(math.Pi)
		l := lerp.NewRemap64(inMin, inMax, outMin, outMax)
		if expected, actual := outMax, l.OutputMaximum(); actual != expected {
			t.Fatalf("Remap64[%v, %v, %v, %v] OutputMaximum: expected[%v] actual[%v]", inMin, inMax, outMin, outMax, expected, actual)
		}
	})
}
