package lerp_test

import (
	"math"
	"testing"

	"github.com/awonak/EuroPiGo/lerp"
)

func TestRemapPoint(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("NewRemapPoint32", func(t *testing.T) {
			in, out := 0, float32(math.Pi)
			if actual := lerp.NewRemapPoint32(in, out); actual == nil {
				t.Fatalf("RemapPoint[%v, %v] NewRemapPoint32: expected[non-nil] actual[nil]", in, out)
			}
		})
		t.Run("NewRemapPoint64", func(t *testing.T) {
			in, out := 0, float32(math.Pi)
			if actual := lerp.NewRemapPoint64(in, out); actual == nil {
				t.Fatalf("RemapPoint[%v, %v] NewRemapPoint64: expected[non-nil] actual[nil]", in, out)
			}
		})
	})

	t.Run("Remap", func(t *testing.T) {
		t.Run("ZeroRange", func(t *testing.T) {
			in, out := 0, float32(math.Pi)
			l := lerp.NewRemapPoint32(in, out)
			if expected, actual := out, l.Remap(in); actual != expected {
				t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
			}
		})
		t.Run("InRange", func(t *testing.T) {
			in, out := 0, float32(math.Pi)
			l := lerp.NewRemapPoint32(in, out)
			if expected, actual := out, l.Remap(in); actual != expected {
				t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
			}
		})

		t.Run("OutOfRange", func(t *testing.T) {
			t.Run("BelowMin", func(t *testing.T) {
				in, out := 0, float32(math.Pi)
				l := lerp.NewRemapPoint32(in, out)
				if expected, actual := out, l.Remap(-2); actual != expected {
					t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				in, out := 0, float32(math.Pi)
				l := lerp.NewRemapPoint32(in, out)
				if expected, actual := out, l.Remap(12); actual != expected {
					t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
				}
			})
		})
	})

	t.Run("Unmap", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			in, out := 0, float32(math.Pi)
			l := lerp.NewRemapPoint32(in, out)
			if expected, actual := in, l.Unmap(out); actual != expected {
				t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
			}
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// Unmap() will work just reply with the "in" point when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				in, out := 0, float32(math.Pi)
				l := lerp.NewRemapPoint32(in, out)
				if expected, actual := in, l.Unmap(out-2); actual != expected {
					t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				in, out := 0, float32(math.Pi)
				l := lerp.NewRemapPoint32(in, out)
				if expected, actual := in, l.Unmap(out+2); actual != expected {
					t.Fatalf("RemapPoint[%v, %v] Remap: expected[%v] actual[%v]", in, out, expected, actual)
				}
			})
		})
	})

	t.Run("MCoeff", func(t *testing.T) {
		in, out := 0, float32(math.Pi)
		l := lerp.NewRemapPoint32(in, out)
		if expected, actual := float32(0.0), l.MCoeff(); actual != expected {
			t.Fatalf("RemapPoint[%v, %v] MCoeff: expected[%v] actual[%v]", in, out, expected, actual)
		}
	})

	t.Run("InputMinimum", func(t *testing.T) {
		in, out := 0, float32(math.Pi)
		l := lerp.NewRemapPoint32(in, out)
		if expected, actual := in, l.InputMinimum(); actual != expected {
			t.Fatalf("RemapPoint[%v, %v] InputMinimum: expected[%v] actual[%v]", in, out, expected, actual)
		}
	})

	t.Run("InputMaximum", func(t *testing.T) {
		in, out := 0, float32(math.Pi)
		l := lerp.NewRemapPoint32(in, out)
		if expected, actual := in, l.InputMaximum(); actual != expected {
			t.Fatalf("RemapPoint[%v, %v] InputMaximum: expected[%v] actual[%v]", in, out, expected, actual)
		}
	})

	t.Run("OutputMinimum", func(t *testing.T) {
		in, out := 0, float32(math.Pi)
		l := lerp.NewRemapPoint32(in, out)
		if expected, actual := out, l.OutputMinimum(); actual != expected {
			t.Fatalf("RemapPoint[%v, %v] OutputMinimum: expected[%v] actual[%v]", in, out, expected, actual)
		}
	})

	t.Run("OutputMaximum", func(t *testing.T) {
		in, out := 0, float32(math.Pi)
		l := lerp.NewRemapPoint32(in, out)
		if expected, actual := out, l.OutputMaximum(); actual != expected {
			t.Fatalf("RemapPoint[%v, %v] OutputMaximum: expected[%v] actual[%v]", in, out, expected, actual)
		}
	})
}
