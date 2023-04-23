package hal

type RandomGenerator interface {
	Configure(config RandomGeneratorConfig) error
}

type RandomGeneratorConfig struct{}
