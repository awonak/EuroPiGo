package module

import (
	"fmt"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ChanceString(chance float32) string {
	return fmt.Sprintf("%3.1f%%", chance*100.0)
}

func ChanceToCV(chance float32) units.CV {
	return units.CV(chance)
}

func CVToChance(cv units.CV) float32 {
	return europim.Clamp(cv.ToFloat32(), 0.0, 1.0)
}
