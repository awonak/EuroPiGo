package hal

// Revision defines an identifier for hardware platforms. See the README.md in the hardware directory for more details.
type Revision int

const (
	RevisionUnknown = Revision(iota)
	Revision0
	Revision1
	Revision2
	// NOTE: always ONLY append to this list, NEVER remove, rename, or reorder
)

// aliases
const (
	EuroPiProto = Revision0
	EuroPi      = Revision1
	EuroPiX     = Revision2
)
