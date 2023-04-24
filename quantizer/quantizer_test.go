package quantizer_test

import (
	"math"
	"testing"

	"github.com/awonak/EuroPiGo/quantizer"
)

func TestNew(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		mode := quantizer.Mode(math.MinInt)
		if actual := quantizer.New[int](mode); actual != nil {
			t.Fatalf("Quantizer[%v] New: expected[nil] actual[non-nil]", mode)
		}
	})
}
