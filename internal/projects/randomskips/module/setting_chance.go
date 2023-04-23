package module

import (
	"fmt"

	"github.com/heucuva/europi/clamp"
	"github.com/heucuva/europi/units"
)

func ChanceString(chance float32) string {
	return fmt.Sprintf("%3.1f%%", chance*100.0)
}

func ChanceToCV(chance float32) units.CV {
	return units.CV(chance)
}

func CVToChance(cv units.CV) float32 {
	return clamp.Clamp(cv.ToFloat32(), 0.0, 1.0)
}
