package math

import "math"

type Lerpable interface {
	~uint8 | ~uint16 | ~int | ~float32 | ~int32 | ~int64
}

func Lerp[V Lerpable](t float32, low, high V) V {
	return V(t*float32(high-low)) + low
}

func LerpRound[V Lerpable](t float32, low, high V) V {
	l := math.Round(float64(t) * float64(high-low))
	return Clamp(V(l)+low, low, high)
}

func InverseLerp[V Lerpable](v, low, high V) float32 {
	if high == low {
		return 0
	}
	return float32(v-low) / float32(high-low)
}
