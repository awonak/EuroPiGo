package rev1

import "github.com/awonak/EuroPiGo/hardware/hal"

type revisionMarker struct{}

var (
	// static check
	_ hal.RevisionMarker = &revisionMarker{}
	// silence linter
	_ = newRevisionMarker
)

func newRevisionMarker() hal.RevisionMarker {
	return &revisionMarker{}
}

// Revision returns the detected revision of the current hardware
func (r *revisionMarker) Revision() hal.Revision {
	return hal.Revision1
}
