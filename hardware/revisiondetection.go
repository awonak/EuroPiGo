package hardware

import (
	"sync"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

// GetRevision returns the currently detected hardware revision.
// If the current revision hasn't been detected, yet, then this call
// will block until it is.
func GetRevision() hal.Revision {
	var waitForDetect sync.WaitGroup
	waitForDetect.Add(1)
	var detectedRevision hal.Revision
	OnRevisionDetected(func(revision hal.Revision) {
		detectedRevision = revision
		waitForDetect.Done()
	})
	waitForDetect.Wait()
	return detectedRevision
}

var (
	onRevisionDetected  chan func(revision hal.Revision)
	revisionChannelInit sync.Once
	revisionWgDone      sync.Once
)

func ensureOnRevisionDetection() {
	revisionChannelInit.Do(func() {
		onRevisionDetected = make(chan func(revision hal.Revision), 10)
	})
}

func OnRevisionDetected(fn func(revision hal.Revision)) {
	if fn == nil {
		return
	}
	ensureOnRevisionDetection()
	onRevisionDetected <- fn
}

// SetDetectedRevision sets the currently detected hardware revision.
// This should not be called directly.
func SetDetectedRevision(opts ...hal.Revision) {
	ensureOnRevisionDetection()
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
