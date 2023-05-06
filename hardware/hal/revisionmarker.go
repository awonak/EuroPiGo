package hal

type RevisionMarker interface {
	Revision() Revision
}

type revisionMarker struct {
	detectedRevision Revision
}

var (
	// static check
	_ RevisionMarker = &revisionMarker{}

	RevisionMark RevisionMarker
)

func NewRevisionMark(opts ...Revision) RevisionMarker {
	r := &revisionMarker{
		detectedRevision: RevisionUnknown,
	}
	if len(opts) > 0 {
		r.detectedRevision = opts[0]
	}
	return r
}

// Revision returns the detected revision of the current hardware
func (r revisionMarker) Revision() Revision {
	return r.detectedRevision
}
