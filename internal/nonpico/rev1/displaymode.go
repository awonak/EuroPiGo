//go:build !pico
// +build !pico

package rev1

type displayMode int

const (
	displayModeSeparate = displayMode(iota)
	displayModeCombined
)
