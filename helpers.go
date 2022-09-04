package europi

func Clamp(value, low, high int) int {
	switch {
	case low < value && value < high:
		return value
	case value > high:
		return high
	case value < low:
		return low
	default:
		return value
	}
}
