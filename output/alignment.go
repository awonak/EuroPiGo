package output

type HorizontalAlignment int

const (
	AlignLeft = HorizontalAlignment(iota)
	AlignCenter
	AlignRight
)

type VerticalAlignment int

const (
	AlignTop = VerticalAlignment(iota)
	AlignMiddle
	AlignBottom
)
