package quantizer

// Quantizer specifies a quantization interface.
type Quantizer[T any] interface {
	QuantizeToIndex(in float32, length int) int
	QuantizeToValue(in float32, list []T) T
}

// New creates a quantizer of the specified `mode`.
func New[T any](mode Mode) Quantizer[T] {
	switch mode {
	case ModeRound:
		return &Round[T]{}
	case ModeTrunc:
		return &Trunc[T]{}
	default:
		// unsupported mode
		return nil
	}
}
