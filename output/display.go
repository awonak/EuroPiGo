package output

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/notoemoji"
	"tinygo.org/x/tinyfont/proggy"
)

const (
	OLEDFreq   = machine.KHz * 400
	OLEDAddr   = ssd1306.Address_128_32
	OLEDWidth  = 128
	OLEDHeight = 32
)

var (
	DefaultChannel   = machine.I2C0
	DefaultFont      = &proggy.TinySZ8pt7b
	DefaultEmojiFont = &notoemoji.NotoEmojiRegular12pt
	White            = color.RGBA{255, 255, 255, 255}
	Black            = color.RGBA{0, 0, 0, 255}
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
	d.WriteLineAligned(text, x, y, AlignLeft, AlignTop)
}

// WriteEmojiLineAligned writes the given emoji text to the display where:
//  x, y is the bottom leftmost pixel of the text
//  alignh is horizontal alignment
//  alignv is vertical alignment
func (d *Display) WriteEmojiLineAligned(text string, x, y int16, alignh HorizontalAlignment, alignv VerticalAlignment) {
	d.writeLineAligned(text, d.font, x, y, alignh, alignv)
}

// WriteLineAligned writes the given text to the display where:
//  x, y is the bottom leftmost pixel of the text
//  alignh is horizontal alignment
//  alignv is vertical alignment
func (d *Display) WriteLineAligned(text string, x, y int16, alignh HorizontalAlignment, alignv VerticalAlignment) {
	d.writeLineAligned(text, d.font, x, y, alignh, alignv)
}

func (d *Display) writeLineAligned(text string, font tinyfont.Fonter, x, y int16, alignh HorizontalAlignment, alignv VerticalAlignment) {
	x0, y0 := x, y
	switch alignh {
	case AlignLeft:
	case AlignCenter:
		_, outerWidth := tinyfont.LineWidth(font, text)
		x0 = (OLEDWidth-int16(outerWidth))/2 - x
	case AlignRight:
		_, outerWidth := tinyfont.LineWidth(font, text)
		x0 = OLEDWidth - int16(outerWidth) - x
	default:
		panic("invalid alignment")
	}
	tinyfont.WriteLine(d, font, x0, y0, text, White)
}

// WriteLineInverse writes the given text to the display in an inverted way where x, y is the bottom leftmost pixel of the text
func (d *Display) WriteLineInverse(text string, x, y int16) {
	d.WriteLineInverseAligned(text, x, y, AlignLeft, AlignTop)
}

// WriteLineInverseAligned writes the given text to the display in an inverted way where:
//  x, y is the bottom leftmost pixel of the text
//  alignh is horizontal alignment
//  alignv is vertical alignment
func (d *Display) WriteLineInverseAligned(text string, x, y int16, alignh HorizontalAlignment, alignv VerticalAlignment) {
	d.writeLineInverseAligned(text, d.font, x, y, alignh, alignv)
}

// WriteEmojiLineInverseAligned writes the given emoji to the display in an inverted way where:
//  x, y is the bottom leftmost pixel of the text
//  alignh is horizontal alignment
//  alignv is vertical alignment
func (d *Display) WriteEmojiLineInverseAligned(text string, x, y int16, alignh HorizontalAlignment, alignv VerticalAlignment) {
	d.writeLineInverseAligned(text, DefaultEmojiFont, x, y, alignh, alignv)
}

func (d *Display) writeLineInverseAligned(text string, font tinyfont.Fonter, x, y int16, alignh HorizontalAlignment, alignv VerticalAlignment) {
	_, outerWidth := tinyfont.LineWidth(font, text)
	outerHeight := int16(font.GetYAdvance())
	x0, y0 := x, y-outerHeight+2
	x1, y1 := x+1, y
	switch alignh {
	case AlignLeft:
	case AlignCenter:
		x0 = (OLEDWidth-int16(outerWidth))/2 - x
		x1 = x0 + 1
	case AlignRight:
		x0 = OLEDWidth - int16(outerWidth) - x
		x1 = x0 + 1
	default:
		panic("invalid alignment")
	}
	switch alignv {
	case AlignTop:
	case AlignMiddle:
		midY := (OLEDHeight - outerHeight) / 2
		y0 += midY
		y1 += midY
	case AlignBottom:
		y1 = OLEDHeight - y1
		y0 = y1 - outerHeight + 2
	default:
		panic("invalid alignment")
	}
	tinydraw.FilledRectangle(d, x0, y0, int16(outerWidth+2), outerHeight, White)
	tinyfont.WriteLine(d, font, x1, y1, text, Black)
}

// DrawHLine draws a horizontal line
func (d *Display) DrawHLine(x, y, xLen int16, c color.RGBA) {
	tinydraw.Line(d, x, y, x+xLen-1, y, c)
}

// DrawLine draws an arbitrary line
func (d *Display) DrawLine(x0, y0, x1, y1 int16, c color.RGBA) {
	tinydraw.Line(d, x0, y0, x1, y1, c)
}
