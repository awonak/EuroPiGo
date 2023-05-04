//go:build !pico
// +build !pico

package europi

func init() {
	DefaultPanicHandler = handlePanicLogger
}
