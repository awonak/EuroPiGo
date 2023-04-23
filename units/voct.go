package units

import "github.com/awonak/EuroPiGo/clamp"

const (
	MinVOct VOct = 0.0
	MaxVOct VOct = 10.0
)

// VOct is a representation of a Volt-per-Octave value
type VOct float32

// ToVolts converts a V/Octave value to a value between 0.0 and 10.0 volts
func (v VOct) ToVolts() float32 {
	return v.ToFloat32()
}

// ToFloat32 returns a V/Octave value to its floating point representation [0.0 .. 10.0]
func (v VOct) ToFloat32() float32 {
	return float32(clamp.Clamp(v, MinVOct, MaxVOct))
}
