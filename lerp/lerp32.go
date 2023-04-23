package lerp

import "github.com/awonak/EuroPiGo/clamp"

type lerp32[T Lerpable] struct {
	b T
	r float32
}

func NewLerp32[T Lerpable](min, max T) Lerper32[T] {
	return lerp32[T]{
		b: min,
		r: float32(max - min),
	}
}

func (l lerp32[T]) Lerp(t float32) T {
	return T(t*l.r) + l.b
}

func (l lerp32[T]) ClampedLerp(t float32) T {
	return clamp.Clamp(T(t*l.r)+l.b, l.b, T(l.r)+l.b)
}

func (l lerp32[T]) LerpRound(t float32) T {
	return T(t*l.r+0.5) + l.b
}

func (l lerp32[T]) ClampedLerpRound(t float32) T {
	return clamp.Clamp(T(t*l.r+0.5)+l.b, l.b, T(l.r)+l.b)
}

func (l lerp32[T]) InverseLerp(v T) float32 {
	if l.r != 0.0 {
		return float32(v-l.b) / l.r
	}
	return 0.0
}

func (l lerp32[T]) ClampedInverseLerp(v T) float32 {
	if l.r != 0.0 {
		return clamp.Clamp(float32(v-l.b)/l.r, 0.0, 1.0)
	}
	return 0.0
}
