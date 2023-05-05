package hardware

import (
	"sync"
	"sync/atomic"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

// GetHardware returns a hardware device based on EuroPi `revision` and hardware `id`.
// a `nil` result means that the hardware was not found or some sort of error occurred.
func GetHardware[T any](revision hal.Revision, id hal.HardwareId) T {
	switch revision {
	case hal.Revision0:
		return rev0.GetHardware[T](id)

	case hal.Revision1:
		return rev1.GetHardware[T](id)

	case hal.Revision2:
		// TODO: implement hardware design of rev2
		//return rev2.GetHardware[T](id)
		fallthrough

	default:
		var none T
		return none
	}
}

// WaitForReady awaits the readiness of the hardware initialization.
// This will block until every aspect of hardware initialization has completed.
func WaitForReady() {
	hardwareReadyCond.L.Lock()
	for {
		ready := hardwareReady.Load()
		if v, ok := ready.(bool); v && ok {
			break
		}
		hardwareReadyCond.Wait()
	}
	hardwareReadyCond.L.Unlock()
}

var (
	hardwareReady     atomic.Value
	hardwareReadyMu   sync.Mutex
	hardwareReadyCond = sync.NewCond(&hardwareReadyMu)
)

// SetReady is used by the hardware initialization code.
// Do not call this function directly.
func SetReady() {
	hardwareReady.Store(true)
	hardwareReadyCond.Broadcast()
}
