package lerp_test

import (
	"testing"

	"github.com/awonak/EuroPiGo/lerp"
)

func TestLerp64(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		min, max := 0, 10
		if actual := lerp.NewLerp64(min, max); actual == nil {
			t.Fatalf("Lerp64[%v, %v] NewLerp64: expected[non-nil] actual[nil]", min, max)
		}
	})

	t.Run("Lerp", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := min, l.Lerp(0.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] Lerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := max, l.Lerp(1.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] Lerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// Lerp() will work as a linear extrapolator when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := -2*max, l.Lerp(-2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] Lerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := 2*max, l.Lerp(2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] Lerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})
	})

	t.Run("ClampedLerp", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := min, l.ClampedLerp(0.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := max, l.ClampedLerp(1.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			t.Run("BelowMin", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := min, l.ClampedLerp(-2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := max, l.ClampedLerp(2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})
	})

	t.Run("LerpRound", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := min, l.LerpRound(0.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] LerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := max, l.LerpRound(1.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] LerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// LerpRound() will work as a linear extrapolator when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := -2*max+1, l.LerpRound(-2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] LerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := 2*max, l.LerpRound(2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] LerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})
	})

	t.Run("ClampedLerpRound", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := min, l.ClampedLerpRound(0.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := max, l.ClampedLerpRound(1.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// ClampedLerpRound() will work as a linear extrapolator when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := min, l.ClampedLerpRound(-2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := max, l.ClampedLerpRound(2.0); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedLerpRound: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})
	})

	t.Run("InverseLerp", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(0.0), l.InverseLerp(min); actual != expected {
					t.Fatalf("Lerp64[%v, %v] InverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Zero", func(t *testing.T) {
				min, max := 5, 5
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(0.0), l.InverseLerp(max); actual != expected {
					t.Fatalf("Lerp64[%v, %v] InverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(1.0), l.InverseLerp(max); actual != expected {
					t.Fatalf("Lerp64[%v, %v] InverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			// InverseLerp() will work as an inverse linear extrapolator when operating out of range
			t.Run("BelowMin", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(-0.2), l.InverseLerp(-2); actual != expected {
					t.Fatalf("Lerp64[%v, %v] InverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(1.2), l.InverseLerp(12); actual != expected {
					t.Fatalf("Lerp64[%v, %v] InverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})
	})

	t.Run("ClampedInverseLerp", func(t *testing.T) {
		t.Run("InRange", func(t *testing.T) {
			t.Run("Min", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(0.0), l.ClampedInverseLerp(min); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedInverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Zero", func(t *testing.T) {
				min, max := 5, 5
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(0.0), l.ClampedInverseLerp(max); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedInverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("Max", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(1.0), l.ClampedInverseLerp(max); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedInverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})

		t.Run("OutOfRange", func(t *testing.T) {
			t.Run("BelowMin", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(0.0), l.ClampedInverseLerp(-2); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedInverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})

			t.Run("AboveMax", func(t *testing.T) {
				min, max := 0, 10
				l := lerp.NewLerp64(min, max)
				if expected, actual := float64(1.0), l.ClampedInverseLerp(12); actual != expected {
					t.Fatalf("Lerp64[%v, %v] ClampedInverseLerp: expected[%v] actual[%v]", min, max, expected, actual)
				}
			})
		})
	})

	t.Run("OutputMinimum", func(t *testing.T) {
		min, max := 0, 10
		l := lerp.NewLerp64(min, max)
		if expected, actual := min, l.OutputMinimum(); actual != expected {
			t.Fatalf("Lerp64[%v, %v] OutputMinimum: expected[%v] actual[%v]", min, max, expected, actual)
		}
	})

	t.Run("OutputMaximum", func(t *testing.T) {
		min, max := 0, 10
		l := lerp.NewLerp64(min, max)
		if expected, actual := max, l.OutputMaximum(); actual != expected {
			t.Fatalf("Lerp64[%v, %v] OutputMaximum: expected[%v] actual[%v]", min, max, expected, actual)
		}
	})
}
