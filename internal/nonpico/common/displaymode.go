//go:build !pico
// +build !pico

package common

type DisplayMode int

const (
	DisplayModeSeparate = DisplayMode(iota)
	DisplayModeCombined
)
