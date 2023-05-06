package units_test

import (
	"testing"
	"time"

	"github.com/awonak/EuroPiGo/units"
)

func TestDurationString(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		zero := time.Duration(0)
		if expected, actual := "0s", units.DurationString(zero); actual != expected {
			t.Fatalf("Duration[%v] DurationString: expected[%s] actual[%s]", zero, expected, actual)
		}
	})

	t.Run("Nanoseconds", func(t *testing.T) {
		ns := time.Duration(time.Nanosecond * 123)
		if expected, actual := "123ns", units.DurationString(ns); actual != expected {
			t.Fatalf("Duration[%v] DurationString: expected[%s] actual[%s]", ns, expected, actual)
		}
	})

	t.Run("Microseconds", func(t *testing.T) {
		us := time.Duration(time.Nanosecond * 12_345)
		if expected, actual := "12.3us", units.DurationString(us); actual != expected {
			t.Fatalf("Duration[%v] DurationString: expected[%s] actual[%s]", us, expected, actual)
		}
	})

	t.Run("Milliseconds", func(t *testing.T) {
		ms := time.Duration(time.Microsecond * 12_345)
		if expected, actual := "12.3ms", units.DurationString(ms); actual != expected {
			t.Fatalf("Duration[%v] DurationString: expected[%s] actual[%s]", ms, expected, actual)
		}
	})

	t.Run("Seconds", func(t *testing.T) {
		s := time.Duration(time.Millisecond * 12_345)
		if expected, actual := "12.3s", units.DurationString(s); actual != expected {
			t.Fatalf("Duration[%v] DurationString: expected[%s] actual[%s]", s, expected, actual)
		}
	})

	t.Run("Minutes", func(t *testing.T) {
		m := time.Duration(time.Millisecond * 754567)
		if expected, actual := "754.6s", units.DurationString(m); actual != expected {
			t.Fatalf("Duration[%v] DurationString: expected[%s] actual[%s]", m, expected, actual)
		}
	})
}
