package europi

// Clamp returns a value that is no lower than `low“ and no higher than `high“.
func Clamp[V int | float32](value, low, high V) V {
	if value > high {
		value = high
	}
	if value < low {
		value = low
	}
	return value
}
