package module

type Config struct {
	Gate   func(high bool)
	Chance float32
}
