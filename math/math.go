package math

type Clampable interface {
	~uint8 | ~uint16 | ~int | ~float32 | ~int32 | ~int64
}

// Clamp returns a value that is no lower than "low" and no higher than "high".
func Clamp[V Clampable](value, low, high V) V {
	if value >= high {
		return high
	}
	if value <= low {
		return low
	}
	return value
}

// Abs returns the absolute value
func Abs(value float32) float32 {
	if value >= 0 {
		return value
	}

	return -value
}
