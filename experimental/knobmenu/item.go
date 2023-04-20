package knobmenu

import "github.com/heucuva/europi/units"

type item struct {
	name     string
	label    string
	stringFn func() string
	updateFn func(value units.CV)
}
