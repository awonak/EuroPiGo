//go:build pico && onscreenpanic
// +build pico,onscreenpanic

package bootstrap

func init() {
	DefaultPanicHandler = handlePanicOnScreenLog
}
