//go:build onscreenpanic
// +build onscreenpanic

package europi

func init() {
	DefaultPanicHandler = handlePanicOnScreenLog
}
