package hardware

import (
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

// GetHardware returns a hardware device based on EuroPi `revision` and hardware `id`.
// a `nil` result means that the hardware was not found or some sort of error occurred.
func GetHardware[T any](revision hal.Revision, id hal.HardwareId) T {
	switch revision {
	case hal.Revision1:
		return rev1.GetHardware[T](id)

	case hal.Revision2:
		// TODO: implement hardware design of rev2
		return rev1.GetHardware[T](id)

	default:
		var none T
		return none
	}
}

// RevisionDetection returns the best (most recent?) match for the hardware installed (or compiled for).
func RevisionDetection() hal.Revision {
	// Iterate in reverse - try to find the newest revision that matches.
	for i := hal.Revision2; i > hal.RevisionUnknown; i-- {
		if rd := GetHardware[hal.RevisionMarker](i, hal.HardwareIdRevisionMarker); rd != nil {
			// use the result of the call - don't just use `i` - in the event there's an alias or redirect involved
			return rd.Revision()
		}
	}
	return hal.RevisionUnknown
}
