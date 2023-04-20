package units

import (
	"fmt"
	"time"
)

type Hertz float32

func (h Hertz) ToPeriod() time.Duration {
	return time.Duration(float32(time.Second) / float32(h))
}

func (h Hertz) String() string {
	switch {
	case h < 0.001:
		return fmt.Sprintf("%3.1fuHz", h*1000000.0)
	case h < 1:
		return fmt.Sprintf("%3.1fmHz", h*1000.0)
	case h >= 1000:
		return fmt.Sprintf("%3.1fkHz", h/1000.0)
	default:
		return fmt.Sprintf("%5.1fHz", h)
	}
}
