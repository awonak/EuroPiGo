//go:build pico
// +build pico

package pico

import (
	"machine"
	"runtime/interrupt"
	"sync"

	"github.com/awonak/EuroPiGo/hardware/rev1"
)

var (
	adcOnce sync.Once
)

type picoAdc struct {
	adc machine.ADC
}

func newPicoAdc(pin machine.Pin) rev1.ADCProvider {
	adcOnce.Do(machine.InitADC)

	adc := &picoAdc{
		adc: machine.ADC{Pin: pin},
	}
	adc.adc.Configure(machine.ADCConfig{})
	return adc
}

func (a *picoAdc) Get(samples int) uint16 {
	if samples == 0 {
		return 0
	}

	var sum int
	state := interrupt.Disable()
	for i := 0; i < samples; i++ {
		sum += int(a.adc.Get())
	}
	interrupt.Restore(state)
	return uint16(sum / samples)
}
