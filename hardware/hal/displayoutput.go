package hal

import "image/color"

type DisplayOutput interface {
	ClearBuffer()
	Size() (x, y int16)
	SetPixel(x, y int16, c color.RGBA)
	Display() error
}

type DisplayOutputConfig struct {
}
