package units

// CV is a normalized representation [0.0 .. 1.0] of a Control Voltage [0.0 .. 5.0] value.
type CV float32

// ToVolts converts a (normalized) CV value to a value between 0.0 and 5.0 volts
func (c CV) ToVolts() float32 {
	v := float32(c)
	range_check(v, 0, 1, "cv")
	return v * 5
}

// ToBipolarCV converts a (normalized) CV value to a (normalized) BipolarCV value
func (c CV) ToBipolarCV() BipolarCV {
	return BipolarCV(c.ToFloat32()*2.0 - 1.0)
}

// ToFloat32 returns a (normalized) CV value to its floating point representation [0.0 .. 1.0]
func (c CV) ToFloat32() float32 {
	return float32(c)
}
