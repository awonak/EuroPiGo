package units_test

import (
	"math"
	"testing"

	"github.com/awonak/EuroPiGo/units"
)

func TestVOctToVolts(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		t.Run("Min", func(t *testing.T) {
			min := units.VOct(0.0)
			if expected, actual := float32(0.0), min.ToVolts(); actual != expected {
				t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", min, expected, actual)
			}
		})

		t.Run("Max", func(t *testing.T) {
			max := units.VOct(10.0)
			if expected, actual := float32(10.0), max.ToVolts(); actual != expected {
				t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", max, expected, actual)
			}
		})
	})
	t.Run("OutOfRange", func(t *testing.T) {
		t.Run("BelowMin", func(t *testing.T) {
			belowMin := units.VOct(-2.0)
			if expected, actual := float32(0.0), belowMin.ToVolts(); actual != expected {
				t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", belowMin, expected, actual)
			}
		})

		t.Run("AboveMax", func(t *testing.T) {
			aboveMax := units.VOct(12.0)
			if expected, actual := float32(10.0), aboveMax.ToVolts(); actual != expected {
				t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", aboveMax, expected, actual)
			}
		})
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.VOct(math.NaN())
		if actual := nan.ToVolts(); !math.IsNaN(float64(actual)) {
			t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", nan, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		t.Run("Neg", func(t *testing.T) {
			negInf := units.VOct(math.Inf(-1))
			if expected, actual := float32(0.0), negInf.ToVolts(); actual != expected {
				t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", negInf, expected, actual)
			}
		})

		t.Run("Pos", func(t *testing.T) {
			posInf := units.VOct(math.Inf(1))
			if expected, actual := float32(10.0), posInf.ToVolts(); actual != expected {
				t.Fatalf("VOct[%v] ToVolts: expected[%f] actual[%f]", posInf, expected, actual)
			}
		})
	})
}

func TestVOctToFloat32(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		t.Run("Min", func(t *testing.T) {
			min := units.VOct(0.0)
			if expected, actual := float32(0.0), min.ToFloat32(); actual != expected {
				t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", min, expected, actual)
			}
		})

		t.Run("Max", func(t *testing.T) {
			max := units.VOct(10.0)
			if expected, actual := float32(10.0), max.ToFloat32(); actual != expected {
				t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", max, expected, actual)
			}
		})
	})
	t.Run("OutOfRange", func(t *testing.T) {
		t.Run("BelowMin", func(t *testing.T) {
			belowMin := units.VOct(-2.0)
			if expected, actual := float32(0.0), belowMin.ToFloat32(); actual != expected {
				t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", belowMin, expected, actual)
			}
		})

		t.Run("BelowMax", func(t *testing.T) {
			aboveMax := units.VOct(122.0)
			if expected, actual := float32(10.0), aboveMax.ToFloat32(); actual != expected {
				t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", aboveMax, expected, actual)
			}
		})
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.VOct(math.NaN())
		if actual := nan.ToFloat32(); !math.IsNaN(float64(actual)) {
			t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", nan, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		t.Run("Neg", func(t *testing.T) {
			negInf := units.VOct(math.Inf(-1))
			if expected, actual := float32(0.0), negInf.ToFloat32(); actual != expected {
				t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", negInf, expected, actual)
			}
		})

		t.Run("Pos", func(t *testing.T) {
			posInf := units.VOct(math.Inf(1))
			if expected, actual := float32(10.0), posInf.ToFloat32(); actual != expected {
				t.Fatalf("VOct[%v] ToFloat32: expected[%f] actual[%f]", posInf, expected, actual)
			}
		})
	})
}
