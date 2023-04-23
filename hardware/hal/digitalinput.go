package hal

import "time"

type DigitalInput interface {
	Configure(config DigitalInputConfig) error
	Value() bool
	Handler(handler func(value bool, deltaTime time.Duration))
	HandlerEx(changes ChangeFlags, handler func(value bool, deltaTime time.Duration))
	HandlerWithDebounce(handler func(value bool, deltaTime time.Duration), delay time.Duration)
}

type DigitalInputConfig struct {
	Samples         uint16
	CalibratedMinAI uint16
	CalibratedMaxAI uint16
}
