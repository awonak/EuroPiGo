package rev1

import (
	"image/color"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

// displayoutput is a wrapper around `ssd1306.Device` for drawing graphics and text to the OLED.
type displayoutput struct {
	dp displayProvider
}

var (
	// static check
	_ hal.DisplayOutput = &displayoutput{}
	// silence linter
	_ = newDisplayOutput
)

type displayProvider interface {
	ClearBuffer()
	Size() (x, y int16)
	SetPixel(x, y int16, c color.RGBA)
	Display() error
}

// newDisplayOutput returns a new Display struct.
func newDisplayOutput(dp displayProvider) hal.DisplayOutput {
	return &displayoutput{
		dp: dp,
	}
}

// Configure updates the device with various configuration parameters
func (d *displayoutput) Configure(config hal.DisplayOutputConfig) error {
	return nil
}

// ClearBuffer clears the internal display buffer for the device
func (d *displayoutput) ClearBuffer() {
	d.dp.ClearBuffer()
}

// Size returns the display resolution for the device
func (d *displayoutput) Size() (x, y int16) {
	return d.dp.Size()
}

// SetPixel sets a specific pixel at coordinates (`x`,`y`) to color `c`.
func (d *displayoutput) SetPixel(x, y int16, c color.RGBA) {
	d.dp.SetPixel(x, y, c)
}

// Display commits the internal buffer to the display device.
// This will update the physical content displayed on the device.
func (d *displayoutput) Display() error {
	return d.dp.Display()
}
