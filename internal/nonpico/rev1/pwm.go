//go:build !pico
// +build !pico

package rev1

import (
	"fmt"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

type nonPicoPwm struct {
	bus event.Bus
	id  hal.HardwareId
	cal envelope.Map[float32, uint16]
	v   float32
}

func newNonPicoPwm(bus event.Bus, id hal.HardwareId) rev1.PWMProvider {
	p := &nonPicoPwm{
		bus: bus,
		id:  id,
		cal: envelope.NewMap32([]envelope.MapEntry[float32, uint16]{
			{
				Input:  rev1.MinOutputVoltage,
				Output: rev1.CalibratedTop,
			},
			{
				Input:  rev1.MaxOutputVoltage,
				Output: rev1.CalibratedOffset,
			},
		}),
	}
	return p
}

func (p *nonPicoPwm) Configure(config hal.VoltageOutputConfig) error {
	return nil
}

func (p *nonPicoPwm) Set(v float32) {
	volts := p.cal.Remap(v)
	p.v = v
	p.bus.Post(fmt.Sprintf("hw_pwm_%d", p.id), HwMessagePwmValue{
		Value: uint16(volts),
	})
}

func (p *nonPicoPwm) Get() float32 {
	return p.v
}

func (p *nonPicoPwm) MinVoltage() float32 {
	return p.cal.InputMinimum()
}

func (p *nonPicoPwm) MaxVoltage() float32 {
	return p.cal.InputMaximum()
}
