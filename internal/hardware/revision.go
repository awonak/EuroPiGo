package hardware

type Revision int

const (
	Revision0 = Revision(iota)
	Revision1
	Revision2
)

// aliases
const (
	EuroPiProto = Revision0
	EuroPi      = Revision1
	EuroPiX     = Revision2
)
