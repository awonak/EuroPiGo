//go:build !pico
// +build !pico

package common

import (
	"fmt"

	"github.com/awonak/EuroPiGo/hardware/common"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/lerp"
)

type nonPicoPwm struct {
	id  hal.HardwareId
	cal lerp.Remapper32[float32, uint16]
	v   float32
}

var (
	// static check
	_ common.PWMProvider = (*nonPicoPwm)(nil)
)

func NewNonPicoPwm(id hal.HardwareId, cal lerp.Remapper32[float32, uint16]) *nonPicoPwm {
	p := &nonPicoPwm{
		id:  id,
		cal: cal,
	}
	return p
}

func (p *nonPicoPwm) Configure(config hal.VoltageOutputConfig) error {
	return nil
}

func (p *nonPicoPwm) Set(v float32) {
	pulseWidth := p.cal.Remap(v)
	p.v = p.cal.Unmap(pulseWidth)
	bus.Post(fmt.Sprintf("hw_pwm_%d", p.id), HwMessagePwmValue{
		Value:   pulseWidth,
		Voltage: p.v,
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
