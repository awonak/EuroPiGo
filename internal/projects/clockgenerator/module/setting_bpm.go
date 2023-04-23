package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

const (
	MinBPM float32 = 0.1
	MaxBPM float32 = 480.0
)

var (
	bpmLerp = lerp.NewLerp32(MinBPM, MaxBPM)
)

func BPMString(bpm float32) string {
	return fmt.Sprintf(`%3.1f`, bpm)
}

func BPMToCV(bpm float32) units.CV {
	return units.CV(bpmLerp.ClampedInverseLerp(bpm))
}

func CVToBPM(cv units.CV) float32 {
	return bpmLerp.ClampedLerpRound(cv.ToFloat32())
}
