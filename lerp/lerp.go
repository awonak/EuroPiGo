package lerp

type Lerpable interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Float interface {
	~float32 | ~float64
}

type Lerper[T Lerpable, F Float] interface {
	Lerp(t F) T
	ClampedLerp(t F) T
	LerpRound(t F) T
	ClampedLerpRound(t F) T
	InverseLerp(v T) F
	ClampedInverseLerp(v T) F
}

type Lerper32[T Lerpable] Lerper[T, float32]

type Lerper64[T Lerpable] Lerper[T, float64]
