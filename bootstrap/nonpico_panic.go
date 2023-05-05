//go:build !pico
// +build !pico

package bootstrap

func init() {
	DefaultPanicHandler = handlePanicLogger
}
