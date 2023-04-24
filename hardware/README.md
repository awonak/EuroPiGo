# hardware

This package is used for obtaining singleton objects for particular hardware, identified by `Revision` and `HardwareId`.

## Revision List

**NOTE**: The full revision list may be found under [hal/revision.go](hal/revision.go).

| Identifier | Alias | Notes |
|----|----|----|
| `Revision0` | `EuroPiProto` | EuroPi prototype developed before revision 1 design was solidified. Due to lack of examples of this hardware, support is limited or non-existent. |
| `Revision1` | `EuroPi` | EuroPi 'production' release, revision 1. |
| `Revision2` | `EuroPiX` | EuroPi X - an improved hardware revision of the EuroPi. Currently in pre-production development hell. |

## Hardware List

**NOTE**: The full hardware list may be found under [hal/hardware.go](hal/hardware.go).

Calling `GetHardware()` with a `Revision` (see above) and `HardwareId` (see below) may return a singleton hardware interface. Check to see if the returned value is `nil` before using it; `nil` is considered to be either a revision detection failure, missing hardware, or some other error.

| Identifier | Alias | Interface | Notes |
|----|----|----|----|
| `HardwareIdInvalid` | | N/A | Always returns a `nil` interface/object |
| `HardwareIdRevisionMarker` | | `hal.RevisionMarker` | Provides an interface to obtain the `Revision` identifier of the currently detected (or compiled-for) hardware. |
| `HardwareIdDigital1Input` | | `hal.DigitalInput` | The Digital Input of the EuroPi. |
| `HardwareIdAnalog1Input` | `HardwareIdAnalogue1Input` | `hal.AnalogInput` | The Analogue Input of the EuroPi. |
| `HardwareIdDisplay1Output` | | `hal.DisplayOutput` | The Display (OLED) of the EuroPi. Provides an interface for determining display resolution, as it might be different between revisions of the EuroPi hardware. |
| `HardwareIdButton1Input` | | `hal.ButtonInput` | The Button 1 gate input of the EuroPi. |
| `HardwareIdButton2Input` | | `hal.ButtonInput` | The Button 2 gate input of the EuroPi. |
| `HardwareIdKnob1Input` | | `hal.KnobInput` | The Knob 1 potentiometer input of the EuroPi. |
| `HardwareIdKnob2Input` | | `hal.KnobInput` | The Knob 2 potentiometer input of the EuroPi. |
| `HardwareIdVoltage1Output` | `HardwareIdCV1Output` | `hal.VoltageOutput` | The #1 `CV` / `V/Octave` output of the EuroPi. While it supports a 0.0 to 10.0 Volts output necessary for `V/Octave` (see `units.VOct`), it can be carefully used with `units.CV` to output a specific range of 0.0 to 5.0 Volts, instead. |
| `HardwareIdVoltage2Output` | `HardwareIdCV2Output` | `hal.VoltageOutput` | The #2 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage3Output` | `HardwareIdCV3Output` | `hal.VoltageOutput` | The #3 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage4Output` | `HardwareIdCV4Output` | `hal.VoltageOutput` | The #4 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage5Output` | `HardwareIdCV5Output` | `hal.VoltageOutput` | The #5 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdVoltage6Output` | `HardwareIdCV6Output` | `hal.VoltageOutput` | The #6 `CV` / `V/Octave` output of the EuroPi. See `HardwareIdVoltage1Output` for more details. |
| `HardwareIdRandom1Generator` | | `hal.RandomGenerator` | Provides an interface to calibrate or seed the random number generator of the hardware. |
