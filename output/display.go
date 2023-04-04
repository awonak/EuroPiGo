package output

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

const (
	OLEDFreq   = machine.KHz * 400
	OLEDAddr   = ssd1306.Address_128_32
	OLEDWidth  = 128
	OLEDHeight = 32
)

var (
	DefaultChannel = machine.I2C0
	DefaultFont    = &proggy.TinySZ8pt7b
	White          = color.RGBA{255, 255, 255, 255}
)

// Display is a wrapper around `ssd1306.Device` for drawing graphics and text to the OLED.
type Display struct {
	ssd1306.Device
	font *tinyfont.Font
}

// NewDisplay returns a new Display struct.
func NewDisplay(channel *machine.I2C, sdaPin, sclPin machine.Pin) *Display {
	channel.Configure(machine.I2CConfig{
		Frequency: OLEDFreq,
		SDA:       sdaPin,
		SCL:       sclPin,
	})

	display := ssd1306.NewI2C(DefaultChannel)
	display.Configure(ssd1306.Config{
		Address: OLEDAddr,
		Width:   OLEDWidth,
		Height:  OLEDHeight,
	})
	return &Display{Device: display, font: DefaultFont}
}

// SetFont overrides the default font used by `WriteLine`.
func (d *Display) SetFont(font *tinyfont.Font) {
	d.font = font
}

// WriteLine writes the given text to the display where x, y is the bottom leftmost pixel of the text.
func (d *Display) WriteLine(text string, x, y int16) {
	tinyfont.WriteLine(d, d.font, x, y, text, White)
}

// DrawHLine draws a horizontal line
func (d *Display) DrawHLine(x, y, xLen int16, c color.RGBA) {
	tinydraw.Line(d, x, y, x+xLen-1, y, c)
}

// DrawLine draws an arbitrary line
func (d *Display) DrawLine(x0, y0, x1, y1 int16, c color.RGBA) {
	tinydraw.Line(d, x0, y0, x1, y1, c)
}
