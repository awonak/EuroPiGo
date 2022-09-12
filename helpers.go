package europi

type Number interface {
	int | float32
}

// Clamp returns a value that is no lower than `low“ and no higher than `high“.
func Clamp[V Number](value, low, high V) V {
	if value > high {
		value = high
	}
	if value < low {
		value = low
	}
	return value
}
