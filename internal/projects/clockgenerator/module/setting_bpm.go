package module

import (
	"fmt"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

const (
	MinBPM float32 = 0.1
	MaxBPM float32 = 480.0
)

func BPMString(bpm float32) string {
	return fmt.Sprintf(`%3.1f`, bpm)
}

func BPMToCV(bpm float32) units.CV {
	return units.CV(europim.InverseLerp(bpm, MinBPM, MaxBPM))
}

func CVToBPM(cv units.CV) float32 {
	return europim.LerpRound(cv.ToFloat32(), MinBPM, MaxBPM)
}
