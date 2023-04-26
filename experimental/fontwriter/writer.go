package fontwriter

import (
	"image/color"

	"github.com/awonak/EuroPiGo/hardware/hal"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

type Writer struct {
	Display hal.DisplayOutput
	Font    tinyfont.Fonter
}

// WriteLine writes the given text to the display where x, y is the bottom leftmost pixel of the text.
func (w *Writer) WriteLine(s string, x, y int16, c color.RGBA) {
	tinyfont.WriteLine(w.Display, w.Font, x, y, s, c)
}

// WriteLineAligned writes the given text to the display where:
//  x, y is the bottom leftmost pixel of the text
//  alignh is horizontal alignment
//  alignv is vertical alignment
func (w *Writer) WriteLineAligned(text string, x, y int16, c color.RGBA, alignh HorizontalAlignment, alignv VerticalAlignment) {
	w.writeLineAligned(text, x, y, c, alignh, alignv)
}

func (w *Writer) writeLineAligned(text string, x, y int16, c color.RGBA, alignh HorizontalAlignment, alignv VerticalAlignment) {
	x0, y0 := x, y
	switch alignh {
	case AlignLeft:
	case AlignCenter:
		dispWidth, _ := w.Display.Size()
		_, outerWidth := tinyfont.LineWidth(w.Font, text)
		x0 = (dispWidth-int16(outerWidth))/2 - x
	case AlignRight:
		dispWidth, _ := w.Display.Size()
		_, outerWidth := tinyfont.LineWidth(w.Font, text)
		x0 = dispWidth - int16(outerWidth) - x
	default:
		panic("invalid alignment")
	}
	tinyfont.WriteLine(w.Display, w.Font, x0, y0, text, c)
}

// WriteLineInverse writes the given text to the display in an inverted way where x, y is the bottom leftmost pixel of the text
func (w *Writer) WriteLineInverse(text string, x, y int16, c color.RGBA) {
	inverseC := color.RGBA{
		R: ^c.R,
		G: ^c.G,
		B: ^c.B,
		A: c.A,
	}
	_, outerWidth := tinyfont.LineWidth(w.Font, text)
	outerHeight := int16(w.Font.GetYAdvance())
	x0, y0 := x, y-outerHeight+2
	x1, y1 := x+1, y
	_ = tinydraw.FilledRectangle(w.Display, x0, y0, int16(outerWidth+2), outerHeight, c)
	tinyfont.WriteLine(w.Display, w.Font, x1, y1, text, inverseC)
}

// WriteLineInverseAligned writes the given text to the display in an inverted way where:
//  x, y is the bottom leftmost pixel of the text
//  alignh is horizontal alignment
//  alignv is vertical alignment
func (w *Writer) WriteLineInverseAligned(text string, x, y int16, c color.RGBA, alignh HorizontalAlignment, alignv VerticalAlignment) {
	w.writeLineInverseAligned(text, w.Font, x, y, c, alignh, alignv)
}

func (w *Writer) writeLineInverseAligned(text string, font tinyfont.Fonter, x, y int16, c color.RGBA, alignh HorizontalAlignment, alignv VerticalAlignment) {
	_, outerWidth := tinyfont.LineWidth(font, text)
	outerHeight := int16(font.GetYAdvance())
	x0, y0 := x, y-outerHeight+2
	x1, y1 := x+1, y
	switch alignh {
	case AlignLeft:
	case AlignCenter:
		dispWidth, _ := w.Display.Size()
		x0 = (dispWidth-int16(outerWidth))/2 - x
		x1 = x0 + 1
	case AlignRight:
		dispWidth, _ := w.Display.Size()
		x0 = dispWidth - int16(outerWidth) - x
		x1 = x0 + 1
	default:
		panic("invalid alignment")
	}
	switch alignv {
	case AlignTop:
	case AlignMiddle:
		_, dispHeight := w.Display.Size()
		midY := (dispHeight - outerHeight) / 2
		y0 += midY
		y1 += midY
	case AlignBottom:
		_, dispHeight := w.Display.Size()
		y1 = dispHeight - y1
		y0 = y1 - outerHeight + 2
	default:
		panic("invalid alignment")
	}

	inverseC := color.RGBA{
		R: ^c.R,
		G: ^c.G,
		B: ^c.B,
		A: c.A,
	}
	_ = tinydraw.FilledRectangle(w.Display, x0, y0, int16(outerWidth+2), outerHeight, c)
	tinyfont.WriteLine(w.Display, w.Font, x1, y1, text, inverseC)
}
