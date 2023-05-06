//go:build pico
// +build pico

package pico

import (
	"machine"

	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev0"
	"github.com/awonak/EuroPiGo/hardware/rev1"
)

// EuroPi Prototype
func initRevision0() {
	rev0.Initialize(rev0.InitializationParameters{
		InputButton1:           newPicoDigitalReader(machine.GPIO15),
		InputButton2:           newPicoDigitalReader(machine.GPIO18),
		InputKnob1:             newPicoAdc(machine.ADC2), // machine.GPIO28
		InputKnob2:             newPicoAdc(machine.ADC1), // machine.GPIO27
		OutputAnalog1:          newPicoPwm(machine.PWM2, machine.GPIO21),
		OutputAnalog2:          newPicoPwm(machine.PWM3, machine.GPIO22),
		OutputAnalog3:          newPicoPwm(machine.PWM1, machine.GPIO19),
		OutputAnalog4:          newPicoPwm(machine.PWM2, machine.GPIO20),
		OutputDigital1:         newPicoPwm(machine.PWM7, machine.GPIO14),
		OutputDigital2:         newPicoPwm(machine.PWM5, machine.GPIO11),
		OutputDigital3:         newPicoPwm(machine.PWM5, machine.GPIO10),
		OutputDigital4:         newPicoPwm(machine.PWM3, machine.GPIO7),
		DeviceRandomGenerator1: &picoRnd{},
	})
}

// EuroPi (original)
func initRevision1() {
	rev1.Initialize(rev1.InitializationParameters{
		InputDigital1:          newPicoDigitalReader(machine.GPIO22),
		InputAnalog1:           newPicoAdc(machine.ADC0), // machine.GPIO26
		OutputDisplay1:         newPicoDisplayOutput(machine.I2C0, machine.GPIO0, machine.GPIO1),
		InputButton1:           newPicoDigitalReader(machine.GPIO4),
		InputButton2:           newPicoDigitalReader(machine.GPIO5),
		InputKnob1:             newPicoAdc(machine.ADC1), // machine.GPIO27
		InputKnob2:             newPicoAdc(machine.ADC2), // machine.GPIO28
		OutputVoltage1:         newPicoPwm(machine.PWM2, machine.GPIO21),
		OutputVoltage2:         newPicoPwm(machine.PWM2, machine.GPIO20),
		OutputVoltage3:         newPicoPwm(machine.PWM0, machine.GPIO16),
		OutputVoltage4:         newPicoPwm(machine.PWM0, machine.GPIO17),
		OutputVoltage5:         newPicoPwm(machine.PWM1, machine.GPIO18),
		OutputVoltage6:         newPicoPwm(machine.PWM1, machine.GPIO19),
		DeviceRandomGenerator1: &picoRnd{},
	})
}

// EuroPi-X
func initRevision2() {
	// TODO: initialize hardware
}

func init() {
	hardware.OnRevisionDetected(func(revision hal.Revision) {
		switch revision {
		case hal.Revision0:
			initRevision0()
		case hal.Revision1:
			initRevision1()
		case hal.Revision2:
			initRevision2()
		default:
		}
		hardware.SetReady()
	})
}
