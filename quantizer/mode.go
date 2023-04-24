package quantizer

// Mode specifies the kind of Quantizer function to be used.
type Mode int

const (
	ModeRound = Mode(iota)
	ModeTrunc
)
