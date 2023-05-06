package rev0

import (
	"context"

	"github.com/awonak/EuroPiGo/hardware/common"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// Pi will be configured during `init()` from platform-specific files.
// See `hardware/pico/pico.go` and `hardware/nonpico/nonpico.go` for more information.
var Pi *EuroPiPrototype

var (
	// static check
	_ hal.Hardware = (*EuroPiPrototype)(nil)
)

type EuroPiPrototype struct {
	common.ContextPi

	// B1 is the Button 1 input on a EuroPi Prototype
	B1 hal.ButtonInput
	// B2 is the Button 2 input on a EuroPi Prototype
	B2 hal.ButtonInput
	// K1 is the Knob 1 input on a EuroPi Prototype
	K1 hal.KnobInput
	// K2 is the Knob 2 input on a EuroPi Prototype
	K2 hal.KnobInput
	// AJ1 is the analog voltage output 1 jack on a EuroPi Prototype. It supports a range of output voltages between 0.0 and 3.3 V.
	AJ1 hal.VoltageOutput
	// AJ2 is the analog voltage output 2 jack on a EuroPi Prototype. It supports a range of output voltages between 0.0 and 3.3 V.
	AJ2 hal.VoltageOutput
	// AJ3 is the analog voltage output 3 jack on a EuroPi Prototype. It supports a range of output voltages between 0.0 and 3.3 V.
	AJ3 hal.VoltageOutput
	// AJ4 is the analog voltage output 4 jack on a EuroPi Prototype. It supports a range of output voltages between 0.0 and 3.3 V.
	AJ4 hal.VoltageOutput
	// DJ1 is the digital voltage output 1 jack on a EuroPi Prototype. It supports output voltages of 0.0 and 3.3 V.
	DJ1 hal.VoltageOutput
	// DJ2 is the digital voltage output 2 jack on a EuroPi Prototype. It supports output voltages of 0.0 and 3.3 V.
	DJ2 hal.VoltageOutput
	// DJ3 is the digital voltage output 3 jack on a EuroPi Prototype. It supports output voltages of 0.0 and 3.3 V.
	DJ3 hal.VoltageOutput
	// DJ4 is the digital voltage output 4 jack on a EuroPi Prototype. It supports output voltages of 0.0 and 3.3 V.
	DJ4 hal.VoltageOutput
	// RND is the random number generator within the EuroPi Prototype.
	RND hal.RandomGenerator
}

func (e *EuroPiPrototype) Context() context.Context {
	return e
}

func (e *EuroPiPrototype) Revision() hal.Revision {
	return hal.Revision0
}

func (e *EuroPiPrototype) Random() hal.RandomGenerator {
	return e.RND
}

func (e *EuroPiPrototype) String() string {
	return "EuroPi Prototype"
}

func (e *EuroPiPrototype) AJ() [4]hal.VoltageOutput {
	return [4]hal.VoltageOutput{e.AJ1, e.AJ2, e.AJ3, e.AJ4}
}

func (e *EuroPiPrototype) DJ() [4]hal.VoltageOutput {
	return [4]hal.VoltageOutput{e.DJ1, e.DJ2, e.DJ3, e.DJ4}
}

func (e *EuroPiPrototype) Button(idx int) hal.ButtonInput {
	switch idx {
	case 0:
		return e.B1
	case 1:
		return e.B2
	default:
		return nil
	}
}

func (e *EuroPiPrototype) Knob(idx int) hal.KnobInput {
	switch idx {
	case 0:
		return e.K1
	case 1:
		return e.K2
	default:
		return nil
	}
}

// GetHardware returns a EuroPi hardware device based on hardware `id`.
// a `nil` result means that the hardware was not found or some sort of error occurred.
func GetHardware[T any](hw hal.HardwareId) T {
	var t T
	if Pi == nil {
		return t
	}

	switch hw {
	case HardwareIdButton1Input:
		t, _ = Pi.B1.(T)
	case HardwareIdButton2Input:
		t, _ = Pi.B2.(T)
	case HardwareIdKnob1Input:
		t, _ = Pi.K1.(T)
	case HardwareIdKnob2Input:
		t, _ = Pi.K2.(T)
	case HardwareIdAnalog1Output:
		t, _ = Pi.AJ1.(T)
	case HardwareIdAnalog2Output:
		t, _ = Pi.AJ2.(T)
	case HardwareIdAnalog3Output:
		t, _ = Pi.AJ3.(T)
	case HardwareIdAnalog4Output:
		t, _ = Pi.AJ4.(T)
	case HardwareIdDigital1Output:
		t, _ = Pi.DJ1.(T)
	case HardwareIdDigital2Output:
		t, _ = Pi.DJ2.(T)
	case HardwareIdDigital3Output:
		t, _ = Pi.DJ3.(T)
	case HardwareIdDigital4Output:
		t, _ = Pi.DJ4.(T)
	case HardwareIdRandom1Generator:
		t, _ = Pi.RND.(T)
	default:
	}
	return t
}

// Initialize sets up the hardware
//
// This is only to be called by the automatic platform initialization functions
func Initialize(params InitializationParameters) {
	Pi = &EuroPiPrototype{
		ContextPi: common.ContextPi{
			Context: context.Background(),
		},
		B1:  common.NewDigitalInput(params.InputButton1),
		B2:  common.NewDigitalInput(params.InputButton2),
		K1:  common.NewAnalogInput(params.InputKnob1, aiInitialConfig),
		K2:  common.NewAnalogInput(params.InputKnob2, aiInitialConfig),
		AJ1: common.NewVoltageOuput(params.OutputAnalog1, cvInitialConfig),
		AJ2: common.NewVoltageOuput(params.OutputAnalog2, cvInitialConfig),
		AJ3: common.NewVoltageOuput(params.OutputAnalog3, cvInitialConfig),
		AJ4: common.NewVoltageOuput(params.OutputAnalog4, cvInitialConfig),
		DJ1: common.NewVoltageOuput(params.OutputDigital1, cvInitialConfig),
		DJ2: common.NewVoltageOuput(params.OutputDigital2, cvInitialConfig),
		DJ3: common.NewVoltageOuput(params.OutputDigital3, cvInitialConfig),
		DJ4: common.NewVoltageOuput(params.OutputDigital4, cvInitialConfig),
		RND: common.NewRandomGenerator(params.DeviceRandomGenerator1),
	}
}

// InitializationParameters is a ferry for hardware functions to the interface layer found here
//
// This is only to be used by the automatic platform initialization functions
type InitializationParameters struct {
	InputButton1           common.DigitalReaderProvider
	InputButton2           common.DigitalReaderProvider
	InputKnob1             common.ADCProvider
	InputKnob2             common.ADCProvider
	OutputAnalog1          common.PWMProvider
	OutputAnalog2          common.PWMProvider
	OutputAnalog3          common.PWMProvider
	OutputAnalog4          common.PWMProvider
	OutputDigital1         common.PWMProvider
	OutputDigital2         common.PWMProvider
	OutputDigital3         common.PWMProvider
	OutputDigital4         common.PWMProvider
	DeviceRandomGenerator1 common.RNDProvider
}
