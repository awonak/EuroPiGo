//go:build europi_rmx

package europi

import "machine"

var (
	DisplayChannel = machine.I2C1
	DisplayI2CSda  = machine.GPIO2
	DisplayI2CScl  = machine.GPIO3
)
