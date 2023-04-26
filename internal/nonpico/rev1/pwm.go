//go:build !pico
// +build !pico

package rev1

import (
	"fmt"
	"math"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

type nonPicoPwm struct {
	bus event.Bus
	id  hal.HardwareId
	v   float32
}

func newNonPicoPwm(bus event.Bus, id hal.HardwareId) rev1.PWMProvider {
	p := &nonPicoPwm{
		bus: bus,
		id:  id,
	}
	return p
}

func (p *nonPicoPwm) Configure(config hal.VoltageOutputConfig) error {
	return nil
}

func (p *nonPicoPwm) Set(v float32, ofs uint16) {
	invertedV := v * math.MaxUint16
	// volts := (float32(o.pwm.Top()) - invertedCv) - o.ofs
	volts := invertedV - float32(ofs)
	p.v = v
	p.bus.Post(fmt.Sprintf("hw_pwm_%d", p.id), HwMessagePwmValue{
		Value: uint16(volts),
	})
}

func (p *nonPicoPwm) Get() float32 {
	return p.v
}
