package units

import "github.com/awonak/EuroPiGo/clamp"

// BipolarCV is a normalized representation [-1.0 .. 1.0] of a Control Voltage [-5.0 .. 5.0] value.
type BipolarCV float32

// ToVolts converts a (normalized) BipolarCV value to a value between -5.0 and 5.0 volts
func (c BipolarCV) ToVolts() float32 {
	return c.ToFloat32() * 5.0
}

// ToCV converts a (normalized) BipolarCV value to a (normalized) CV value and a sign bit
func (c BipolarCV) ToCV() (cv CV, sign int) {
	if c < 0.0 {
		return CV(-c.ToFloat32()), -1
	}
	return CV(c.ToFloat32()), 1
}

// ToFloat32 returns a (normalized) BipolarCV value to its floating point representation [-1.0 .. 1.0]
func (c BipolarCV) ToFloat32() float32 {
	return clamp.Clamp(float32(c), -1.0, 1.0)
}
