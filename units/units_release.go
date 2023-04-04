//go:build !debug
// +build !debug

package units

func range_check[T ~float32 | ~float64](v, min, max T, kind string) {
}
