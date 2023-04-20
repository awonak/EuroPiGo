# Random Skips

A random gate skipper based on YouTube video observations made about the operation of the Ladik S-090 module.

## Scope of This App

The scope of this app is to drive the CV 1 output as a gate output based on a percentage chance of 33%. When the input gate (or internal clock gate) goes high (CV >= 0.8V), then a random value is generated and compared against the chance that's provided - if the probability is sufficient enough, then the gate is let through for as long as it is still high on the input. The moment the gate goes low, the output also goes low and the detection process starts again.

### Inputs

- Digital Input = clock input (optional, see below)

### Outputs

- CV 1 = Random Gate output

## Using Random Skips

### Changing Screens

Long-pressing (>=650ms) Button 2 on the EuroPi will transition to the next display in the chain. If you transition past the last item in the display chain, then the display will cycle to the first item.

The order of the displays is:
- Main display
- Random Skips configuration
- Performance clock configuration

#### Main Display

The main display shows the voltages of the CV outputs on the EuroPi as well as the enabled status of the internal performance clock.

While Random Skips is operating, you can toggle between using the external clock (default mode at startup) and the internal clock by pressing Button 1 on the EuroPi while on the main screen. When the internal clock mode is active, you will be informed by seeing a small bar ( `_` ) in the upper-left corner of the display.

#### Random Skips Configuration

By default, the settings of Random Skips are:
- Chance: 50.0%

When on the Random Skips Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

#### Performance Clock Configuration

By default, the settings of the Performance Clock are:
- Clock Rate: 120.0 BPM
- Gate Duration: 100.0 ms

When on the Performance Clock Configuration screen, pressing Button 1 on the EuroPi will cycle through the configuration items. The currently selected item for edit will be identified by an asterisk (`*`) character and it may be updated by turning Knob 1 of the EuroPi. Updates are applied immediately.

## Special Thanks

- Adam Wonak
- Charlotte Cox
- Allen Synthesis
- Ladik.eu
- Mouser Electronics
- Waveshare Electronics
- Raspberry Pi Foundation
