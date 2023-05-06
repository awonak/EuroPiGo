package units

import (
	"fmt"
	"time"
)

type Hertz float32

func (h Hertz) ToPeriod() time.Duration {
	if h == 0 {
		return 0
	}
	return time.Duration(float32(time.Second) / float32(h))
}

func (h Hertz) String() string {
	switch {
	case h < 0.000_001:
		return fmt.Sprintf("%3.1fnHz", h*1_000_000_000.0)
	case h < 0.001:
		return fmt.Sprintf("%3.1fuHz", h*1_000_000.0)
	case h < 1.0:
		return fmt.Sprintf("%3.1fmHz", h*1_000.0)
	case h < 1_000.0:
		return fmt.Sprintf("%3.1fHz", h)
	case h < 1_000_000.0:
		return fmt.Sprintf("%3.1fkHz", h/1_000.0)
	case h < 1_000_000_000.0:
		return fmt.Sprintf("%3.1fMHz", h/1_000_000.0)
	case h < 1_000_000_000_000.0:
		return fmt.Sprintf("%3.1fGHz", h/1_000_000_000.0)
	default:
		// use scientific notation
		return fmt.Sprintf("%3.1gHz", h)
	}
}
