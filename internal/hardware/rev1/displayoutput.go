package rev1

import (
	"image/color"

	"github.com/heucuva/europi/internal/hardware/hal"
)

// displayoutput is a wrapper around `ssd1306.Device` for drawing graphics and text to the OLED.
type displayoutput struct {
	dp displayProvider
}

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

func (d *displayoutput) ClearBuffer() {
	d.dp.ClearBuffer()
}

func (d *displayoutput) Size() (x, y int16) {
	return d.dp.Size()
}
func (d *displayoutput) SetPixel(x, y int16, c color.RGBA) {
	d.dp.SetPixel(x, y, c)
}

func (d *displayoutput) Display() error {
	return d.dp.Display()
}
