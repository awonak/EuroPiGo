package module

import (
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

const (
	MinGateDuration time.Duration = time.Microsecond
	MaxGateDuration time.Duration = time.Millisecond * 990
)

func GateDurationString(dur time.Duration) string {
	return units.DurationString(dur)
}

func GateDurationToCV(dur time.Duration) units.CV {
	return units.CV(europim.InverseLerp(dur, MinGateDuration, MaxGateDuration))
}

func CVToGateDuration(cv units.CV) time.Duration {
	return europim.Lerp[time.Duration](cv.ToFloat32(), MinGateDuration, MaxGateDuration)
}
