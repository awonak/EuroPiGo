package output

import (
	"machine"
)

// PWM is an interface for interacting with a machine.pwmGroup
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	SetTop(top uint32)
	Get(channel uint8) uint32
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}
