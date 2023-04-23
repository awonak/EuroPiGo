package hardware

import "github.com/awonak/EuroPiGo/hardware/hal"

type Revision = hal.Revision

const (
	RevisionUnknown = hal.RevisionUnknown
	Revision0       = hal.Revision0
	Revision1       = hal.Revision1
	Revision2       = hal.Revision2
)

// aliases
const (
	EuroPiProto = Revision0
	EuroPi      = Revision1
	EuroPiX     = Revision2
)
