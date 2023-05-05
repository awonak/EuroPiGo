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

var (
	onRevisionDetected                                    = make(chan func(revision hal.Revision), 10)
	OnRevisionDetected chan<- func(revision hal.Revision) = onRevisionDetected
	revisionWgDone     sync.Once
	hardwareReady      atomic.Value
	hardwareReadyMu    sync.Mutex
	hardwareReadyCond  = sync.NewCond(&hardwareReadyMu)
)

func SetDetectedRevision(opts ...hal.Revision) {
	// need to be sure it's ready before we can done() it
	hal.RevisionMark = hal.NewRevisionMark(opts...)
	revisionWgDone.Do(func() {
		go func() {
			for fn := range onRevisionDetected {
				if fn != nil {
					fn(hal.RevisionMark.Revision())
				}
			}
		}()
	})
}

func SetReady() {
	hardwareReady.Store(true)
	hardwareReadyCond.Broadcast()
}

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

func GetRevision() hal.Revision {
	var waitForDetect sync.WaitGroup
	waitForDetect.Add(1)
	var detectedRevision hal.Revision
	OnRevisionDetected <- func(revision hal.Revision) {
		detectedRevision = revision
		waitForDetect.Done()
	}
	waitForDetect.Wait()
	return detectedRevision
}
