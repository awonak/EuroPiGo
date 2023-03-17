package math

type Lerpable interface {
	~uint8 | ~uint16 | ~int | ~float32
}

func Lerp[V Lerpable](t float32, low, high V) V {
	return V(t*float32(high-low)) + low
}
