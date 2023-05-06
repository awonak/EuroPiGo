package rev1

import (
	"context"

	"github.com/awonak/EuroPiGo/hardware/common"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// Pi will be configured during `init()` from platform-specific files.
// See `hardware/pico/pico.go` and `hardware/nonpico/nonpico.go` for more information.
var Pi *EuroPi

type EuroPi struct {
	common.ContextPi

	// DI is the Digital Input on a EuroPi
	DI hal.DigitalInput
	// AI is the Analogue Input on a EuroPi
	AI hal.AnalogInput
	// OLED is the display output on a EuroPi
	OLED hal.DisplayOutput
	// B1 is the Button 1 input on a EuroPi
	B1 hal.ButtonInput
	// B2 is the Button 2 input on a EuroPi
	B2 hal.ButtonInput
	// K1 is the Knob 1 input on a EuroPi
	K1 hal.KnobInput
	// K2 is the Knob 2 input on a EuroPi
	K2 hal.KnobInput
	// CV1 is the voltage output 1 jack on a EuroPi. It supports a range of output voltages between 0.0 and 10.0 V.
	CV1 hal.VoltageOutput
	// CV2 is the voltage output 2 jack on a EuroPi. It supports a range of output voltages between 0.0 and 10.0 V.
	CV2 hal.VoltageOutput
	// CV3 is the voltage output 3 jack on a EuroPi. It supports a range of output voltages between 0.0 and 10.0 V.
	CV3 hal.VoltageOutput
	// CV4 is the voltage output 4 jack on a EuroPi. It supports a range of output voltages between 0.0 and 10.0 V.
	CV4 hal.VoltageOutput
	// CV5 is the voltage output 5 jack on a EuroPi. It supports a range of output voltages between 0.0 and 10.0 V.
	CV5 hal.VoltageOutput
	// CV6 is the voltage output 6 jack on a EuroPi. It supports a range of output voltages between 0.0 and 10.0 V.
	CV6 hal.VoltageOutput
	// RND is the random number generator within the EuroPi
	RND hal.RandomGenerator
}

func (e *EuroPi) Context() context.Context {
	return e
}

func (e *EuroPi) Revision() hal.Revision {
	return hal.Revision1
}

func (e *EuroPi) Random() hal.RandomGenerator {
	return e.RND
}

func (e *EuroPi) String() string {
	return "EuroPi"
}

func (e *EuroPi) CV() [6]hal.VoltageOutput {
	return [6]hal.VoltageOutput{e.CV1, e.CV2, e.CV3, e.CV4, e.CV5, e.CV6}
}

func (e *EuroPi) Button(idx int) hal.ButtonInput {
	switch idx {
	case 0:
		return e.B1
	case 1:
		return e.B2
	default:
		return nil
	}
}

func (e *EuroPi) Knob(idx int) hal.KnobInput {
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
	case HardwareIdDigital1Input:
		t, _ = Pi.DI.(T)
	case HardwareIdAnalog1Input:
		t, _ = Pi.AI.(T)
	case HardwareIdDisplay1Output:
		t, _ = Pi.OLED.(T)
	case HardwareIdButton1Input:
		t, _ = Pi.B1.(T)
	case HardwareIdButton2Input:
		t, _ = Pi.B2.(T)
	case HardwareIdKnob1Input:
		t, _ = Pi.K1.(T)
	case HardwareIdKnob2Input:
		t, _ = Pi.K2.(T)
	case HardwareIdCV1Output:
		t, _ = Pi.CV1.(T)
	case HardwareIdCV2Output:
		t, _ = Pi.CV2.(T)
	case HardwareIdCV3Output:
		t, _ = Pi.CV3.(T)
	case HardwareIdCV4Output:
		t, _ = Pi.CV4.(T)
	case HardwareIdCV5Output:
		t, _ = Pi.CV5.(T)
	case HardwareIdCV6Output:
		t, _ = Pi.CV6.(T)
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
	Pi = &EuroPi{
		ContextPi: common.ContextPi{
			Context: context.Background(),
		},
		DI:   common.NewDigitalInput(params.InputDigital1),
		AI:   common.NewAnalogInput(params.InputAnalog1, aiInitialConfig),
		OLED: common.NewDisplayOutput(params.OutputDisplay1),
		B1:   common.NewDigitalInput(params.InputButton1),
		B2:   common.NewDigitalInput(params.InputButton2),
		K1:   common.NewAnalogInput(params.InputKnob1, aiInitialConfig),
		K2:   common.NewAnalogInput(params.InputKnob2, aiInitialConfig),
		CV1:  common.NewVoltageOuput(params.OutputVoltage1, cvInitialConfig),
		CV2:  common.NewVoltageOuput(params.OutputVoltage2, cvInitialConfig),
		CV3:  common.NewVoltageOuput(params.OutputVoltage3, cvInitialConfig),
		CV4:  common.NewVoltageOuput(params.OutputVoltage4, cvInitialConfig),
		CV5:  common.NewVoltageOuput(params.OutputVoltage5, cvInitialConfig),
		CV6:  common.NewVoltageOuput(params.OutputVoltage6, cvInitialConfig),
		RND:  common.NewRandomGenerator(params.DeviceRandomGenerator1),
	}
}

// InitializationParameters is a ferry for hardware functions to the interface layer found here
//
// This is only to be used by the automatic platform initialization functions
type InitializationParameters struct {
	InputDigital1          common.DigitalReaderProvider
	InputAnalog1           common.ADCProvider
	OutputDisplay1         common.DisplayProvider
	InputButton1           common.DigitalReaderProvider
	InputButton2           common.DigitalReaderProvider
	InputKnob1             common.ADCProvider
	InputKnob2             common.ADCProvider
	OutputVoltage1         common.PWMProvider
	OutputVoltage2         common.PWMProvider
	OutputVoltage3         common.PWMProvider
	OutputVoltage4         common.PWMProvider
	OutputVoltage5         common.PWMProvider
	OutputVoltage6         common.PWMProvider
	DeviceRandomGenerator1 common.RNDProvider
}
