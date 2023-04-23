package hardware

import (
	"github.com/heucuva/europi/internal/hardware/hal"
	"github.com/heucuva/europi/internal/hardware/rev1"
)

func GetHardware[T any](revision Revision, id hal.HardwareId) T {
	switch revision {
	case Revision1:
		return rev1.GetHardware[T](id)

	case Revision2:
		// TODO: implement hardware design of rev2
		return rev1.GetHardware[T](id)

	default:
		var none T
		return none
	}
}

func RevisionDetection() Revision {
	for i := Revision0; i <= Revision2; i++ {
		if rd := GetHardware[hal.RevisionMarker](i, hal.HardwareIdRevisionMarker); rd != nil {
			// use the result of the call - don't just use `i` - in the event there's an alias or redirect involved
			return rd.Revision()
		}
	}
	return RevisionUnknown
}
