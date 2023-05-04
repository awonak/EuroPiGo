//go:build pico && onscreenpanic
// +build pico,onscreenpanic

package europi

func init() {
	DefaultPanicHandler = handlePanicOnScreenLog
}
