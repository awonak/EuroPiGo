//go:build pico
// +build pico

package pico

import (
	"fmt"
	"machine"
	"math"
	"runtime/interrupt"
	"runtime/volatile"
	"time"

	"github.com/awonak/EuroPiGo/experimental/envelope"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
)

type picoPwm struct {
	pwm      pwmGroup
	pin      machine.Pin
	ch       uint8
	v        uint32
	period   time.Duration
	wavefold bool
	cal      envelope.Map[float32, uint16]
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

const (
	picoPwmModeAnalogRevision0 = picoPwmMode(iota)
	picoPwmModeDigitalRevision0
	picoPwmModeAnalogRevision1
)

func newPicoPwm(pwm pwmGroup, pin machine.Pin, mode picoPwmMode) *picoPwm {
	p := &picoPwm{
		pwm:    pwm,
		pin:    pin,
		period: rev0.DefaultPWMPeriod,
		cal: envelope.NewMap32([]envelope.MapEntry[float32, uint16]{
			{
				Input:  rev0.MinOutputVoltage,
				Output: rev0.CalibratedTop,
			},
			{
				Input:  rev0.MaxOutputVoltage,
				Output: rev0.CalibratedOffset,
			},
		}),
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

	p.pwm.SetTop(uint32(p.cal.OutputMaximum()))
	ch, err := p.pwm.Channel(p.pin)
	if err != nil {
		return fmt.Errorf("pwm Channel error: %w", err)
	}
	p.ch = ch

	p.wavefold = config.PerformWavefold

	return nil
}

func (p *picoPwm) Set(v float32) {
	if p.wavefold {
		if v < 0.0 {
			v = -v
		}
	}
	volts := p.cal.Remap(v)
	state := interrupt.Disable()
	p.pwm.Set(p.ch, uint32(volts))
	interrupt.Restore(state)
	volatile.StoreUint32(&p.v, math.Float32bits(v))
}

func (p *picoPwm) Get() float32 {
	return math.Float32frombits(volatile.LoadUint32(&p.v))
}

func (p *picoPwm) MinVoltage() float32 {
	return p.cal.InputMinimum()
}

func (p *picoPwm) MaxVoltage() float32 {
	return p.cal.InputMaximum()
}
