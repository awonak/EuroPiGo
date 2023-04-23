package hal

type ChangeFlags int

const (
	ChangeRising = ChangeFlags(1 << iota)
	ChangeFalling
)

const (
	ChangeNone = ChangeFlags(0)
	ChangeAny  = ChangeRising | ChangeFalling
)
