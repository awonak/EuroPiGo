//go:build pico
// +build pico

package pico

import (
	"fmt"
	"machine"
	"math"
	"runtime/interrupt"
	"sync/atomic"
	"time"

	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
)

type picoPwm struct {
	pwm       pwmGroup
	pin       machine.Pin
	ch        uint8
	v         uint32
	period    time.Duration
	monopolar bool
	cal       envelope.Map[float32, uint16]
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

type picoPwmMode int

func newPicoPwm(pwm pwmGroup, pin machine.Pin) *picoPwm {
	p := &picoPwm{
		pwm:    pwm,
		pin:    pin,
		period: rev0.DefaultPWMPeriod,
		// NOTE: cal must be set non-nil by Configure() at least 1 time
	}
	return p
}

func (p *picoPwm) Configure(config hal.VoltageOutputConfig) error {
	state := interrupt.Disable()
	defer interrupt.Restore(state)

	if config.Period != 0 {
		p.period = config.Period
	}

	err := p.pwm.Configure(machine.PWMConfig{
		Period: uint64(p.period.Nanoseconds()),
	})
	if err != nil {
		return fmt.Errorf("pwm Configure error: %w", err)
	}

	if config.Calibration != nil {
		p.cal = config.Calibration
	}

	if any(p.cal) == nil {
		return fmt.Errorf("pwm Configure error: Calibration must be non-nil")
	}

	p.pwm.SetTop(uint32(p.cal.OutputMaximum()))
	ch, err := p.pwm.Channel(p.pin)
	if err != nil {
		return fmt.Errorf("pwm Channel error: %w", err)
	}
	p.ch = ch

	p.monopolar = config.Monopolar

	return nil
}

func (p *picoPwm) Set(v float32) {
	if p.monopolar {
		if v < 0.0 {
			v = -v
		}
	}
	volts := p.cal.Remap(v)
	state := interrupt.Disable()
	p.pwm.Set(p.ch, uint32(volts))
	interrupt.Restore(state)
	atomic.StoreUint32(&p.v, math.Float32bits(v))
}

func (p *picoPwm) Get() float32 {
	return math.Float32frombits(atomic.LoadUint32(&p.v))
}

func (p *picoPwm) MinVoltage() float32 {
	return p.cal.InputMinimum()
}

func (p *picoPwm) MaxVoltage() float32 {
	return p.cal.InputMaximum()
}
