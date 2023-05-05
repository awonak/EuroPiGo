//go:build !pico && websim
// +build !pico,websim

package bootstrap

func init() {
	defaultWebSimEnabled = true
}
