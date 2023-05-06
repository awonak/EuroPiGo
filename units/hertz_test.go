package units_test

import (
	"testing"
	"time"

	"github.com/awonak/EuroPiGo/units"
)

func TestHertzToPeriod(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		zero := units.Hertz(0)
		if expected, actual := time.Duration(0), zero.ToPeriod(); actual != expected {
			t.Fatalf("Hertz[%v] ToPeriod: expected[%s] actual[%s]", zero, expected, actual)
		}
	})

	t.Run("Seconds", func(t *testing.T) {
		s := units.Hertz(0.1)
		if expected, actual := time.Second*10, s.ToPeriod(); actual != expected {
			t.Fatalf("Hertz[%v] ToPeriod: expected[%s] actual[%s]", s, expected, actual)
		}
	})

	t.Run("Milliseconds", func(t *testing.T) {
		ms := units.Hertz(100.0)
		if expected, actual := time.Millisecond*10, ms.ToPeriod(); actual != expected {
			t.Fatalf("Hertz[%v] ToPeriod: expected[%s] actual[%s]", ms, expected, actual)
		}
	})

	t.Run("Microseconds", func(t *testing.T) {
		us := units.Hertz(100_000.0)
		if expected, actual := time.Microsecond*10, us.ToPeriod(); actual != expected {
			t.Fatalf("Hertz[%v] ToPeriod: expected[%s] actual[%s]", us, expected, actual)
		}
	})

	t.Run("Nanoseconds", func(t *testing.T) {
		ns := units.Hertz(100_000_000.0)
		if expected, actual := time.Nanosecond*10, ns.ToPeriod(); actual != expected {
			t.Fatalf("Hertz[%v] ToPeriod: expected[%s] actual[%s]", ns, expected, actual)
		}
	})
}

func TestHertzString(t *testing.T) {
	t.Run("NanoHertz", func(t *testing.T) {
		nhz := units.Hertz(0.000_000_1)
		if expected, actual := "100.0nHz", nhz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", nhz, expected, actual)
		}
	})

	t.Run("MicroHertz", func(t *testing.T) {
		uhz := units.Hertz(0.000_1)
		if expected, actual := "100.0uHz", uhz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", uhz, expected, actual)
		}
	})

	t.Run("MilliHertz", func(t *testing.T) {
		mhz := units.Hertz(0.1)
		if expected, actual := "100.0mHz", mhz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", mhz, expected, actual)
		}
	})

	t.Run("Hertz", func(t *testing.T) {
		hz := units.Hertz(100.0)
		if expected, actual := "100.0Hz", hz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", hz, expected, actual)
		}
	})

	t.Run("KiloHertz", func(t *testing.T) {
		khz := units.Hertz(100_000.0)
		if expected, actual := "100.0kHz", khz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", khz, expected, actual)
		}
	})

	t.Run("MegaHertz", func(t *testing.T) {
		mhz := units.Hertz(100_000_000.0)
		if expected, actual := "100.0MHz", mhz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", mhz, expected, actual)
		}
	})

	t.Run("GigaHertz", func(t *testing.T) {
		ghz := units.Hertz(100_000_000_000.0)
		if expected, actual := "100.0GHz", ghz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", ghz, expected, actual)
		}
	})

	t.Run("Other", func(t *testing.T) {
		hz := units.Hertz(100_000_000_000_000.0)
		if expected, actual := "1e+14Hz", hz.String(); actual != expected {
			t.Fatalf("Hertz[%v] String: expected[%s] actual[%s]", hz, expected, actual)
		}
	})
}
