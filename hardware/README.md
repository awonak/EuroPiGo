# hardware

This package is used for obtaining singleton objects for particular hardware, identified by `Revision` and `HardwareId`.

## Revision List

**NOTE**: The full revision list may be found under [hal/revision.go](hal/revision.go).

| Identifier | Alias | (non-pico) Build Flags | Notes |
|----|----|----|----|
| `Revision0` | `EuroPiProto` | | EuroPi prototype developed before revision 1 design was solidified. Due to lack of examples of this hardware, support is limited or non-existent. |
| `Revision1` | `EuroPi` | `revision1` or `europi` | EuroPi 'production' release, revision 1. |
| `Revision2` | `EuroPiX` | `revision2` or `europix` | EuroPi X - an improved hardware revision of the EuroPi. Currently in pre-production development hell. |

## Hardware List

**NOTE**: The full hardware list may be found under [hal/hardware.go](hal/hardware.go).

Calling `GetHardware()` with a `Revision` (see above) and `HardwareId` (see below) may return a singleton hardware interface. Check to see if the returned value is `nil` before using it; `nil` is considered to be either a revision detection failure, missing hardware, or some other error.

| HardwareId | HardwareId Alias | Interface | EuroPi Prototype | EuroPi | EuroPi-X | Notes |
|----|----|----|----|----|----|----|
| `HardwareIdInvalid` | | N/A | | | | Always returns a `nil` interface/object |
| `HardwareIdRevisionMarker` | | `hal.RevisionMarker` | | | | Provides an interface to obtain the `Revision` identifier of the currently detected (or compiled-for) hardware. |
| `HardwareIdDigital1Input` | | `hal.DigitalInput` | | `InputDigital1` | | The Digital Input of the EuroPi. |
| `HardwareIdAnalog1Input` | `HardwareIdAnalogue1Input` | `hal.AnalogInput` | | `InputAnalog1` | | The Analogue Input of the EuroPi. |
| `HardwareIdDisplay1Output` | | `hal.DisplayOutput` | | `OutputDisplay1` | | The Display (OLED) of the EuroPi. Provides an interface for determining display resolution, as it might be different between revisions of the EuroPi hardware. |
| `HardwareIdButton1Input` | | `hal.ButtonInput` | `InputButton1` | `InputButton1` | | The Button 1 gate input of the EuroPi. |
| `HardwareIdButton2Input` | | `hal.ButtonInput` | `InputButton2` | `InputButton2` | | The Button 2 gate input of the EuroPi. |
| `HardwareIdKnob1Input` | | `hal.KnobInput` | `InputKnob1` | `InputKnob1` | | The Knob 1 potentiometer input of the EuroPi. |
| `HardwareIdKnob2Input` | | `hal.KnobInput` | `InputKnob2` | `InputKnob2` | | The Knob 2 potentiometer input of the EuroPi. |
| `HardwareIdVoltage1Output` | | `hal.VoltageOutput` | `OutputAnalog1` | `OutputVoltage1` | | The #1 `CV` / `V/Octave` output of the EuroPi. While the EuroPi supports a 0.0 to 10.0 Volts output necessary for `V/Octave` (see `units.VOct`), it can be carefully used with `units.CV` to output a specific range of 0.0 to 5.0 Volts, instead. For the EuroPi Prototype, the range is 0.0 to 3.3 Volts. |
| `HardwareIdVoltage2Output` | | `hal.VoltageOutput` | `OutputAnalog2` | `OutputVoltage2` | | The #2 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage3Output` | | `hal.VoltageOutput` | `OutputAnalog3` | `OutputVoltage3` | | The #3 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage4Output` | | `hal.VoltageOutput` | `OutputAnalog4` | `OutputVoltage4` | | The #4 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage5Output` | | `hal.VoltageOutput` | `OutputDigital1` | `OutputVoltage5` | | The #5 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage6Output` | | `hal.VoltageOutput` | `OutputDigital2` | `OutputVoltage6` | | The #6 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdRandom1Generator` | | `hal.RandomGenerator` | `DeviceRandomGenerator1` | `DeviceRandomGenerator1` | | Provides an interface to calibrate or seed the random number generator of the hardware. |
| `HardwareIdVoltage7Output` | | `hal.VoltageOutput` | `OutputDigital3` | | | The #7 `CV` / `V/Octave` output of the EuroPi Prototype. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage8Output` | | `hal.VoltageOutput` | `OutputDigital4` | | | The #8 `CV` / `V/Octave` output of the EuroPi Prototype. See `HardwareIdVoltage1Output` for more details. |
