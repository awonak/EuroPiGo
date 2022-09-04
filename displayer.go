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

var White = color.RGBA{255, 255, 255, 255}

type Displayer interface {
	WriteLine(text string, x, y int16)
}

type Display struct {
	ssd1306.Device
}

func NewDisplay(channel *machine.I2C, sdaPin, sclPin machine.Pin) *Display {
	channel.Configure(machine.I2CConfig{
		Frequency: OLEDFreq,
		SDA:       sdaPin,
		SCL:       sclPin,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: OLEDAddr,
		Width:   OLEDWidth,
		Height:  OLEDHeight,
	})
	return &Display{display}
}

func (d *Display) WriteLine(text string, x, y int16) {
	tinyfont.WriteLine(d, &proggy.TinySZ8pt7b, x, y, text, White)
}
