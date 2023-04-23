package units_test

import (
	"math"
	"testing"

	"github.com/awonak/EuroPiGo/units"
)

func TestBipolarCVToVolts(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		min := units.BipolarCV(-1.0)
		if expected, actual := float32(-5.0), min.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", min, expected, actual)
		}

		zero := units.BipolarCV(0.0)
		if expected, actual := float32(0.0), zero.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", zero, expected, actual)
		}

		max := units.BipolarCV(1.0)
		if expected, actual := float32(5.0), max.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", max, expected, actual)
		}
	})
	t.Run("OutOfRange", func(t *testing.T) {
		belowMin := units.BipolarCV(-2.0)
		if expected, actual := float32(-5.0), belowMin.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", belowMin, expected, actual)
		}

		aboveMax := units.BipolarCV(2.0)
		if expected, actual := float32(5.0), aboveMax.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", aboveMax, expected, actual)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.BipolarCV(math.NaN())
		if actual := nan.ToVolts(); !math.IsNaN(float64(actual)) {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", nan, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		negInf := units.BipolarCV(math.Inf(-1))
		if expected, actual := float32(-5.0), negInf.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", negInf, expected, actual)
		}

		posInf := units.BipolarCV(math.Inf(1))
		if expected, actual := float32(5.0), posInf.ToVolts(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToVolts: expected[%f] actual[%f]", posInf, expected, actual)
		}
	})
}

func TestBipolarCVToCV(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		min := units.BipolarCV(-1.0)
		expectedMin, expectedMinSign := units.CV(1.0), -1
		if actual, actualSign := min.ToCV(); actual != expectedMin || actualSign != expectedMinSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", min, expectedMin, expectedMinSign, actual, actualSign)
		}

		zero := units.BipolarCV(0.0)
		expectedZero, expectedZeroSign := units.CV(0.0), 1
		if actual, actualSign := zero.ToCV(); actual != expectedZero || actualSign != expectedZeroSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", zero, expectedZero, expectedZeroSign, actual, actualSign)
		}

		max := units.BipolarCV(1.0)
		expectedMax, expectedMaxSign := units.CV(1.0), 1
		if actual, actualSign := max.ToCV(); actual != expectedMax || actualSign != expectedMaxSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", max, expectedMax, expectedMaxSign, actual, actualSign)
		}
	})
	t.Run("OutOfRange", func(t *testing.T) {
		belowMin := units.BipolarCV(-2.0)
		expectedBelowMin, expectedBelowMinSign := units.CV(1.0), -1
		if actual, actualSign := belowMin.ToCV(); actual != expectedBelowMin || actualSign != expectedBelowMinSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", belowMin, expectedBelowMin, expectedBelowMinSign, actual, actualSign)
		}

		aboveMax := units.BipolarCV(2.0)
		expectedAboveMax, expectedAboveMaxSign := units.CV(1.0), 1
		if actual, actualSign := aboveMax.ToCV(); actual != expectedAboveMax || actualSign != expectedAboveMaxSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", aboveMax, expectedAboveMax, expectedAboveMaxSign, actual, actualSign)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.BipolarCV(math.NaN())
		expectedNanSign := 1
		if actual, actualSign := nan.ToCV(); !math.IsNaN(float64(actual)) || actualSign != expectedNanSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", nan, math.NaN(), expectedNanSign, actual, actualSign)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		negInf := units.BipolarCV(math.Inf(-1))
		expectedNegInf, expectedNegInfSign := units.CV(1.0), -1
		if actual, actualSign := negInf.ToCV(); actual != expectedNegInf || actualSign != expectedNegInfSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", negInf, expectedNegInf, expectedNegInfSign, actual, actualSign)
		}

		posInf := units.BipolarCV(math.Inf(1))
		expectedPosInf, expectedPosInfSign := units.CV(1.0), 1
		if actual, actualSign := posInf.ToCV(); actual != expectedPosInf || actualSign != expectedPosInfSign {
			t.Fatalf("BipolarCV[%v] ToCV: expected[%f sign(%d)] actual[%f sign(%d)]", posInf, expectedPosInf, expectedPosInfSign, actual, actualSign)
		}
	})
}

func TestBipolarCVToFloat32(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		min := units.BipolarCV(-1.0)
		if expected, actual := float32(-1.0), min.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", min, expected, actual)
		}

		zero := units.BipolarCV(0.0)
		if expected, actual := float32(0.0), zero.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", zero, expected, actual)
		}

		max := units.BipolarCV(1.0)
		if expected, actual := float32(1.0), max.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", max, expected, actual)
		}
	})
	t.Run("OutOfRange", func(t *testing.T) {
		belowMin := units.BipolarCV(-2.0)
		if expected, actual := float32(-1.0), belowMin.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", belowMin, expected, actual)
		}

		aboveMax := units.BipolarCV(2.0)
		if expected, actual := float32(1.0), aboveMax.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", aboveMax, expected, actual)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.BipolarCV(math.NaN())
		if actual := nan.ToFloat32(); !math.IsNaN(float64(actual)) {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", nan, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		negInf := units.BipolarCV(math.Inf(-1))
		if expected, actual := float32(-1.0), negInf.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", negInf, expected, actual)
		}

		posInf := units.BipolarCV(math.Inf(1))
		if expected, actual := float32(1.0), posInf.ToFloat32(); actual != expected {
			t.Fatalf("BipolarCV[%v] ToFloat32: expected[%f] actual[%f]", posInf, expected, actual)
		}
	})
}
