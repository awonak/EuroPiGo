package lerp

import "github.com/heucuva/europi/clamp"

type lerp64[T Lerpable] struct {
	b T
	r float64
}

func NewLerp64[T Lerpable](min, max T) Lerper64[T] {
	return lerp64[T]{
		b: min,
		r: float64(max - min),
	}
}

func (l lerp64[T]) Lerp(t float64) T {
	return T(t*l.r) + l.b
}

func (l lerp64[T]) ClampedLerp(t float64) T {
	return clamp.Clamp(T(t*l.r)+l.b, l.b, T(l.r)+l.b)
}

func (l lerp64[T]) LerpRound(t float64) T {
	return T(t*l.r+0.5) + l.b
}

func (l lerp64[T]) ClampedLerpRound(t float64) T {
	return clamp.Clamp(T(t*l.r+0.5)+l.b, l.b, T(l.r)+l.b)
}

func (l lerp64[T]) InverseLerp(v T) float64 {
	if l.r != 0.0 {
		return float64(v-l.b) / l.r
	}
	return 0.0
}

func (l lerp64[T]) ClampedInverseLerp(v T) float64 {
	if l.r != 0.0 {
		return clamp.Clamp(float64(v-l.b)/l.r, 0.0, 1.0)
	}
	return 0.0
}
