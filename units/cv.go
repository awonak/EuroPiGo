package units

import "github.com/awonak/EuroPiGo/clamp"

// CV is a normalized representation [0.0 .. 1.0] of a Control Voltage [0.0 .. 5.0] value.
type CV float32

// ToVolts converts a (normalized) CV value to a value between 0.0 and 5.0 volts
func (c CV) ToVolts() float32 {
	return c.ToFloat32() * 5.0
}

// ToBipolarCV converts a (normalized) CV value to a (normalized) BipolarCV value
func (c CV) ToBipolarCV(sign int) BipolarCV {
	bc := BipolarCV(c.ToFloat32())
	if sign < 0 {
		return -bc
	}
	return bc
}

// ToFloat32 returns a (normalized) CV value to its floating point representation [0.0 .. 1.0]
func (c CV) ToFloat32() float32 {
	return clamp.Clamp(float32(c), 0.0, 1.0)
}
