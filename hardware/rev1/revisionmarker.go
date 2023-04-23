package rev1

import "github.com/awonak/EuroPiGo/hardware/hal"

type revisionMarker struct{}

func newRevisionMarker() hal.RevisionMarker {
	return &revisionMarker{}
}

func (r *revisionMarker) Revision() hal.Revision {
	return hal.Revision1
}
