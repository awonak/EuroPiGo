package module

import (
	"time"

	"github.com/heucuva/europi/lerp"
	"github.com/heucuva/europi/units"
)

const (
	MinGateDuration time.Duration = time.Microsecond
	MaxGateDuration time.Duration = time.Millisecond * 990
)

var (
	gateDurationLerp = lerp.NewLerp32(MinGateDuration, MaxGateDuration)
)

func GateDurationString(dur time.Duration) string {
	return units.DurationString(dur)
}

func GateDurationToCV(dur time.Duration) units.CV {
	return units.CV(gateDurationLerp.ClampedInverseLerp(dur))
}

func CVToGateDuration(cv units.CV) time.Duration {
	return gateDurationLerp.ClampedLerpRound(cv.ToFloat32())
}
