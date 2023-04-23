package hal

type Revision int

const (
	RevisionUnknown = Revision(iota)
	Revision0
	Revision1
	Revision2
)

type RevisionMarker interface {
	Revision() Revision
}
