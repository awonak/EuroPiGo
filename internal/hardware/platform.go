package hardware

import (
	"github.com/heucuva/europi/internal/hardware/hal"
	"github.com/heucuva/europi/internal/hardware/rev1"
)

func GetHardware[T any](revision Revision, hw hal.HardwareId) T {
	switch revision {
	case Revision1:
		hw, _ := rev1.GetHardware(hw).(T)
		return hw

	case Revision2:
		// TODO: implement hardware design of rev2
		hw, _ := rev1.GetHardware(hw).(T)
		return hw

	default:
		var none T
		return none
	}
}
