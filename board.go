//go:build europi

package europi

import "machine"

var (
	DisplayChannel = machine.I2C0
	DisplayI2CSda  = machine.GPIO0
	DisplayI2CScl  = machine.GPIO1
)
