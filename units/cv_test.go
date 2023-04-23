package units_test

import (
	"math"
	"testing"

	"github.com/awonak/EuroPiGo/units"
)

func TestCVToVolts(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		min := units.CV(0.0)
		if expected, actual := float32(0.0), min.ToVolts(); actual != expected {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", min, expected, actual)
		}

		max := units.CV(1.0)
		if expected, actual := float32(5.0), max.ToVolts(); actual != expected {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", max, expected, actual)
		}
	})
	t.Run("OutOfRange", func(t *testing.T) {
		belowMin := units.CV(-2.0)
		if expected, actual := float32(0.0), belowMin.ToVolts(); actual != expected {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", belowMin, expected, actual)
		}

		aboveMax := units.CV(2.0)
		if expected, actual := float32(5.0), aboveMax.ToVolts(); actual != expected {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", aboveMax, expected, actual)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.CV(math.NaN())
		if actual := nan.ToVolts(); !math.IsNaN(float64(actual)) {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", nan, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		negInf := units.CV(math.Inf(-1))
		if expected, actual := float32(0.0), negInf.ToVolts(); actual != expected {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", negInf, expected, actual)
		}

		posInf := units.CV(math.Inf(1))
		if expected, actual := float32(5.0), posInf.ToVolts(); actual != expected {
			t.Fatalf("CV[%v] ToVolts: expected[%f] actual[%f]", posInf, expected, actual)
		}
	})
}

func TestCVToBipolarCV(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		min, minSign := units.CV(1.0), -1
		if actual, expected := min.ToBipolarCV(minSign), units.BipolarCV(-1.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", min, minSign, expected, actual)
		}

		zero, zeroSign := units.CV(0.0), 1
		if actual, expected := zero.ToBipolarCV(zeroSign), units.BipolarCV(0.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", zero, zeroSign, expected, actual)
		}

		max, maxSign := units.CV(1.0), 1
		if actual, expected := max.ToBipolarCV(maxSign), units.BipolarCV(1.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", max, maxSign, expected, actual)
		}
	})
	t.Run("OutOfRange", func(t *testing.T) {
		belowMin, belowMinSign := units.CV(2.0), -1
		if actual, expected := belowMin.ToBipolarCV(belowMinSign), units.BipolarCV(-1.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", belowMin, belowMinSign, expected, actual)
		}

		aboveMax, aboveMaxSign := units.CV(2.0), 1
		if actual, expected := aboveMax.ToBipolarCV(aboveMaxSign), units.BipolarCV(1.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", aboveMax, aboveMaxSign, expected, actual)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		nan, nanSign := units.CV(math.NaN()), 1
		if actual := nan.ToBipolarCV(nanSign); !math.IsNaN(float64(actual)) {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", nan, nanSign, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		negInf, negInfSign := units.CV(math.Inf(1)), -1
		if actual, expected := negInf.ToBipolarCV(negInfSign), units.BipolarCV(-1.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", negInf, negInfSign, expected, actual)
		}

		posInf, posInfSign := units.CV(math.Inf(1)), 1
		if actual, expected := posInf.ToBipolarCV(posInfSign), units.BipolarCV(1.0); actual != expected {
			t.Fatalf("CV[%v sign(%d)] ToBipolarCV: expected[%v] actual[%v]", posInf, posInfSign, expected, actual)
		}
	})
}

func TestCVToFloat32(t *testing.T) {
	t.Run("InRange", func(t *testing.T) {
		min := units.CV(0.0)
		if expected, actual := float32(0.0), min.ToFloat32(); actual != expected {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", min, expected, actual)
		}

		max := units.CV(1.0)
		if expected, actual := float32(1.0), max.ToFloat32(); actual != expected {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", max, expected, actual)
		}
	})
	t.Run("OutOfRange", func(t *testing.T) {
		belowMin := units.CV(-2.0)
		if expected, actual := float32(0.0), belowMin.ToFloat32(); actual != expected {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", belowMin, expected, actual)
		}

		aboveMax := units.CV(2.0)
		if expected, actual := float32(1.0), aboveMax.ToFloat32(); actual != expected {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", aboveMax, expected, actual)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		nan := units.CV(math.NaN())
		if actual := nan.ToFloat32(); !math.IsNaN(float64(actual)) {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", nan, math.NaN(), actual)
		}
	})

	t.Run("Inf", func(t *testing.T) {
		negInf := units.CV(math.Inf(-1))
		if expected, actual := float32(0.0), negInf.ToFloat32(); actual != expected {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", negInf, expected, actual)
		}

		posInf := units.CV(math.Inf(1))
		if expected, actual := float32(1.0), posInf.ToFloat32(); actual != expected {
			t.Fatalf("CV[%v] ToFloat32: expected[%f] actual[%f]", posInf, expected, actual)
		}
	})
}
