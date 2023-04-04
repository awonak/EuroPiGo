//go:build debug
// +build debug

package units

import (
	"fmt"
)

func range_check[T ~float32 | ~float64](v, min, max T, kind string) {
	if v < min || v > max {
		panic(fmt.Errorf("%w: %v", fmt.Errorf("%s out of range", kind), v))
	}
}
