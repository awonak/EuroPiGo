package knobmenu

import "github.com/awonak/EuroPiGo/units"

type item struct {
	name     string
	label    string
	stringFn func() string
	updateFn func(value units.CV)
}
