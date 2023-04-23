package clamp

type Clampable interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Clamp[T Clampable](v, min, max T) T {
	if v > max {
		v = max
	}
	if v < min {
		v = min
	}
	return v
}
