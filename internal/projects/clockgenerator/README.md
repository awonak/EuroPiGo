# Clock Generator

A simple gate/clock generator based on YouTube video observations made about the operation of the ALM Pamela's NEW Workout module.

## Scope of This App

The scope of this app is to drive the CV 1 output as a gate output.

### Outputs

- CV 1 = Gate output

## Using Clock Generator

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Clock Generator configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi as well as the enabled status of the Clock Generator.

While Clock Generator is operating, you can toggle its activation mode (default mode at startup is `on`) by pressing Button 1 on the EuroPi while on the main screen. When the clock is active, you will be informed by seeing a small bar ( `_` ) in the upper-left corner of the display.

#### Clock Generator Configuration

By default, the settings of Clock Generator are:
- BPM: 120.0
- Gate Duration: 100.0 ms


When on the Clock Generator Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- ALM
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
