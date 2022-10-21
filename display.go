package europi

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

const (
	OLEDFreq   = machine.TWI_FREQ_400KHZ
	OLEDAddr   = 0x3C
	OLEDWidth  = 128
	OLEDHeight = 32
)

var (
	DefaultFont = &proggy.TinySZ8pt7b
	White       = color.RGBA{255, 255, 255, 255}
)

type display struct {
	ssd1306.Device

	font *tinyfont.Font
}

func newDisplay(channel *machine.I2C, sdaPin, sclPin machine.Pin) *display {
	channel.Configure(machine.I2CConfig{
		Frequency: OLEDFreq,
		SDA:       sdaPin,
		SCL:       sclPin,
	})

	d := ssd1306.NewI2C(channel)
	d.Configure(ssd1306.Config{
		Address: OLEDAddr,
		Width:   OLEDWidth,
		Height:  OLEDHeight,
	})
	return &display{Device: d, font: DefaultFont}
}

// Font overrides the default font used by `WriteLine`.
func (d *display) Font(font *tinyfont.Font) {
	d.font = font
}

// WriteLine writes the given text to the display where x, y is the bottom leftmost pixel of the text.
func (d *display) WriteLine(text string, x, y int16) {
	tinyfont.WriteLine(d, d.font, x, y, text, White)
}
