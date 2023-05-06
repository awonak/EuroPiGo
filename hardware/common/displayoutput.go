package common

import (
	"image/color"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

// DisplayOutput is a wrapper around `ssd1306.Device` for drawing graphics and text to the OLED.
type DisplayOutput struct {
	dp DisplayProvider
}

var (
	// static check
	_ hal.DisplayOutput = (*DisplayOutput)(nil)
	// silence linter
	_ = NewDisplayOutput
)

type DisplayProvider interface {
	ClearBuffer()
	Size() (x, y int16)
	SetPixel(x, y int16, c color.RGBA)
	Display() error
}

// NewDisplayOutput returns a new Display struct.
func NewDisplayOutput(dp DisplayProvider) *DisplayOutput {
	if dp == nil {
		return nil
	}
	return &DisplayOutput{
		dp: dp,
	}
}

// Configure updates the device with various configuration parameters
func (d *DisplayOutput) Configure(config hal.DisplayOutputConfig) error {
	return nil
}

// ClearBuffer clears the internal display buffer for the device
func (d *DisplayOutput) ClearBuffer() {
	d.dp.ClearBuffer()
}

// Size returns the display resolution for the device
func (d *DisplayOutput) Size() (x, y int16) {
	return d.dp.Size()
}

// SetPixel sets a specific pixel at coordinates (`x`,`y`) to color `c`.
func (d *DisplayOutput) SetPixel(x, y int16, c color.RGBA) {
	d.dp.SetPixel(x, y, c)
}

// Display commits the internal buffer to the display device.
// This will update the physical content displayed on the device.
func (d *DisplayOutput) Display() error {
	return d.dp.Display()
}
