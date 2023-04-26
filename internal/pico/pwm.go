//go:build pico
// +build pico

package pico

import (
	"fmt"
	"machine"
	"math"
	"runtime/interrupt"
	"runtime/volatile"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

type picoPwm struct {
	pwm pwmGroup
	pin machine.Pin
	ch  uint8
	v   uint32
}

// pwmGroup is an interface for interacting with a machine.pwmGroup
type pwmGroup interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	SetTop(top uint32)
	Get(channel uint8) uint32
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

func newPicoPwm(pwm pwmGroup, pin machine.Pin) rev1.PWMProvider {
	p := &picoPwm{
		pwm: pwm,
		pin: pin,
	}
	return p
}

func (p *picoPwm) Configure(config hal.VoltageOutputConfig) error {
	state := interrupt.Disable()
	defer interrupt.Restore(state)

	err := p.pwm.Configure(machine.PWMConfig{
		Period: uint64(config.Period.Nanoseconds()),
	})
	if err != nil {
		return fmt.Errorf("pwm Configure error: %w", err)
	}

	p.pwm.SetTop(uint32(config.Top))
	ch, err := p.pwm.Channel(p.pin)
	if err != nil {
		return fmt.Errorf("pwm Channel error: %w", err)
	}
	p.ch = ch

	return nil
}

func (p *picoPwm) Set(v float32, ofs uint16) {
	invertedV := v * float32(p.pwm.Top())
	// volts := (float32(o.pwm.Top()) - invertedCv) - o.ofs
	volts := invertedV - float32(ofs)
	state := interrupt.Disable()
	p.pwm.Set(p.ch, uint32(volts))
	interrupt.Restore(state)
	volatile.StoreUint32(&p.v, math.Float32bits(v))
}

func (p *picoPwm) Get() float32 {
	return math.Float32frombits(volatile.LoadUint32(&p.v))
}
