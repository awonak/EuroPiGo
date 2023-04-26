//go:build pico
// +build pico

package pico

import (
	"image/color"
	"machine"

	"github.com/awonak/EuroPiGo/hardware/rev1"
	"tinygo.org/x/drivers/ssd1306"
)

const (
	oledFreq   = machine.KHz * 400
	oledAddr   = ssd1306.Address_128_32
	oledWidth  = 128
	oledHeight = 32
)

type picoDisplayOutput struct {
	dev ssd1306.Device
}

func newPicoDisplayOutput(channel *machine.I2C, sdaPin, sclPin machine.Pin) rev1.DisplayProvider {
	channel.Configure(machine.I2CConfig{
		Frequency: oledFreq,
		SDA:       sdaPin,
		SCL:       sclPin,
	})

	display := ssd1306.NewI2C(channel)
	display.Configure(ssd1306.Config{
		Address: oledAddr,
		Width:   oledWidth,
		Height:  oledHeight,
	})

	dp := &picoDisplayOutput{
		dev: display,
	}

	return dp
}

func (d *picoDisplayOutput) ClearBuffer() {
	d.dev.ClearBuffer()
}

func (d *picoDisplayOutput) Size() (x, y int16) {
	return d.dev.Size()
}
func (d *picoDisplayOutput) SetPixel(x, y int16, c color.RGBA) {
	d.dev.SetPixel(x, y, c)
}

func (d *picoDisplayOutput) Display() error {
	return d.dev.Display()
}
