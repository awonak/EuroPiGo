package units

import (
	"fmt"
	"time"
)

func DurationString(dur time.Duration) string {
	switch {
	case dur < time.Millisecond:
		return fmt.Sprintf("%3.1fus", dur.Seconds()*1000000.0)
	case dur < time.Second:
		return fmt.Sprintf("%3.1fms", dur.Seconds()*1000.0)
	default:
		return fmt.Sprint(dur)
	}
}
