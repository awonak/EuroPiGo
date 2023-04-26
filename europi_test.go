package europi_test

import (
	"testing"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

func TestNew(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		// unless we're running on a pico, we should expect a nil back for the default revision
		if actual := europi.New(); actual != nil {
			t.Fatal("EuroPi New: expected[nil] actual[non-nil]")
		}
	})

	t.Run("Revision1", func(t *testing.T) {
		if actual := europi.NewFrom(hal.Revision1); actual == nil {
			t.Fatal("EuroPi New: expected[non-nil] actual[nil]")
		}
	})
}
