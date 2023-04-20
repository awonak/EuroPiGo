package module

import "time"

const (
	DefaultGateDuration = time.Millisecond * 100
)

type Config struct {
	BPM          float32
	GateDuration time.Duration
	Enabled      bool
	ClockOut     func(high bool)
}
