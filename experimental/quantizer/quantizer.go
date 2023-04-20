package quantizer

type Quantizer[T any] interface {
	QuantizeToIndex(in float32, length int) int
	QuantizeToValue(in float32, list []T) T
}
