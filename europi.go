package europi // import "github.com/awonak/EuroPiGo"

import (
	"github.com/awonak/EuroPiGo/hardware"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// EuroPi is the collection of component wrappers used to interact with the module.
type EuroPi struct {
	Display hal.DisplayOutput
	DI      hal.DigitalInput
	AI      hal.AnalogInput
	B1      hal.ButtonInput
	B2      hal.ButtonInput
	K1      hal.KnobInput
	K2      hal.KnobInput
	CV1     hal.VoltageOutput
	CV2     hal.VoltageOutput
	CV3     hal.VoltageOutput
	CV4     hal.VoltageOutput
	CV5     hal.VoltageOutput
	CV6     hal.VoltageOutput
	CV      [6]hal.VoltageOutput
	RND     hal.RandomGenerator
}

// New will return a new EuroPi struct.
func New(opts ...hal.Revision) *EuroPi {
	var revision hal.Revision
	if len(opts) > 0 {
		revision = opts[0]
	} else {
		// attempt to detect hardware revision
		revision = hardware.RevisionDetection()
	}

	if revision == hal.RevisionUnknown {
		// could not detect revision
		return nil
	}

	cv1 := hardware.GetHardware[hal.VoltageOutput](revision, hal.HardwareIdVoltage1Output)
	cv2 := hardware.GetHardware[hal.VoltageOutput](revision, hal.HardwareIdVoltage2Output)
	cv3 := hardware.GetHardware[hal.VoltageOutput](revision, hal.HardwareIdVoltage3Output)
	cv4 := hardware.GetHardware[hal.VoltageOutput](revision, hal.HardwareIdVoltage4Output)
	cv5 := hardware.GetHardware[hal.VoltageOutput](revision, hal.HardwareIdVoltage5Output)
	cv6 := hardware.GetHardware[hal.VoltageOutput](revision, hal.HardwareIdVoltage6Output)

	e := &EuroPi{
		Display: hardware.GetHardware[hal.DisplayOutput](revision, hal.HardwareIdDisplay1Output),

		DI: hardware.GetHardware[hal.DigitalInput](revision, hal.HardwareIdDigital1Input),
		AI: hardware.GetHardware[hal.AnalogInput](revision, hal.HardwareIdAnalog1Input),

		B1: hardware.GetHardware[hal.ButtonInput](revision, hal.HardwareIdButton1Input),
		B2: hardware.GetHardware[hal.ButtonInput](revision, hal.HardwareIdButton2Input),

		K1: hardware.GetHardware[hal.KnobInput](revision, hal.HardwareIdKnob1Input),
		K2: hardware.GetHardware[hal.KnobInput](revision, hal.HardwareIdKnob2Input),

		CV1: cv1,
		CV2: cv2,
		CV3: cv3,
		CV4: cv4,
		CV5: cv5,
		CV6: cv5,
		CV:  [6]hal.VoltageOutput{cv1, cv2, cv3, cv4, cv5, cv6},
		RND: hardware.GetHardware[hal.RandomGenerator](revision, hal.HardwareIdRandom1Generator),
	}

	return e
}
