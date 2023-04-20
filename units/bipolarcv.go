package units

// BipolarCV is a normalized representation [-1.0 .. 1.0] of a Control Voltage [-5.0 .. 5.0] value.
type BipolarCV float32

// ToVolts converts a (normalized) BipolarCV value to a value between -5.0 and 5.0 volts
func (c BipolarCV) ToVolts() float32 {
	v := float32(c)
	range_check(v, -1, 1, "bipolarcv")
	return v * 5
}

// ToCV converts a (normalized) BipolarCV value to a (normalized) CV value
func (c BipolarCV) ToCV() CV {
	return CV((c.ToFloat32() + 1.0) * 0.5)
}

// ToFloat32 returns a (normalized) BipolarCV value to its floating point representation [-1.0 .. 1.0]
func (c BipolarCV) ToFloat32() float32 {
	return float32(c)
}
