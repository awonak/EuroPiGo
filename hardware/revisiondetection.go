package hardware

import (
	"sync"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

// GetRevision returns the currently detected hardware revision.
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

var (
	onRevisionDetected                                    = make(chan func(revision hal.Revision), 10)
	OnRevisionDetected chan<- func(revision hal.Revision) = onRevisionDetected
	revisionWgDone     sync.Once
)

// SetDetectedRevision sets the currently detected hardware revision.
// This should not be called directly.
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
