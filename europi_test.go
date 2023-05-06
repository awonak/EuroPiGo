package europi_test

import (
	"testing"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

func TestNew(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		// unless we're running on a pico, we should expect a nil back for the default revision
		hardware.SetDetectedRevision(hal.RevisionUnknown)
		if actual := europi.New(); actual != nil {
			t.Fatalf("EuroPi New: expected[nil] actual[%T]", actual)
		}
	})

	t.Run("Revision0", func(t *testing.T) {
		hardware.SetDetectedRevision(hal.Revision0)
		pi := europi.New()
		switch actual := pi.(type) {
		case *rev0.EuroPiPrototype:
			if actual == nil {
				t.Fatal("EuroPi New: expected[non-nil] actual[nil]")
			}
		case nil:
			t.Fatal("EuroPi New: expected[non-nil] actual[nil]")
		default:
			t.Fatalf("EuroPi New: expected[EuroPi Prototype] actual[%v]", actual)
		}
	})

	t.Run("Revision1", func(t *testing.T) {
		hardware.SetDetectedRevision(hal.Revision1)
		pi := europi.New()
		switch actual := pi.(type) {
		case *rev1.EuroPi:
			if actual == nil {
				t.Fatal("EuroPi New: expected[non-nil] actual[nil]")
			}
		case nil:
			t.Fatal("EuroPi New: expected[non-nil] actual[nil]")
		default:
			t.Fatalf("EuroPi New: expected[EuroPi] actual[%v]", actual)
		}
	})
}

func TestNewFrom(t *testing.T) {
	t.Run("Revision0", func(t *testing.T) {
		hardware.SetDetectedRevision(hal.Revision0)
		if actual, _ := europi.NewFrom(hal.Revision0).(*rev0.EuroPiPrototype); actual == nil {
			t.Fatalf("EuroPi NewFrom: expected[EuroPiPrototype] actual[%T]", actual)
		}
	})

	t.Run("Revision1", func(t *testing.T) {
		hardware.SetDetectedRevision(hal.Revision1)
		if actual, _ := europi.NewFrom(hal.Revision1).(*rev1.EuroPi); actual == nil {
			t.Fatalf("EuroPi NewFrom: expected[EuroPi] actual[%T]", actual)
		}
	})
}
