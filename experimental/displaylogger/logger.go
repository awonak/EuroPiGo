package displaylogger

import (
	"io"
	"strings"

	"github.com/heucuva/europi/experimental/draw"
	"github.com/heucuva/europi/experimental/fontwriter"
	"github.com/heucuva/europi/internal/hardware/hal"
	"tinygo.org/x/tinyfont/proggy"
)

var (
	DefaultFont = &proggy.TinySZ8pt7b
)

type Logger interface {
	io.Writer
	Flush()
}

type logger struct {
	sb      strings.Builder
	display hal.DisplayOutput
	writer  fontwriter.Writer
}

func NewLogger(display hal.DisplayOutput) Logger {
	return &logger{
		sb:      strings.Builder{},
		display: display,
		writer: fontwriter.Writer{
			Display: display,
			Font:    DefaultFont,
		},
	}
}

func (w *logger) Write(p []byte) (n int, err error) {
	n, err = w.sb.Write(p)
	if err != nil {
		return
	}

	w.repaint()
	return
}

func (w *logger) Flush() {
	w.repaint()
}

func (w *logger) repaint() {
	str := w.sb.String()

	w.display.ClearBuffer()

	lines := strings.Split(str, "\n")
	w.sb.Reset()
	_, maxY := w.display.Size()
	yAdv := w.writer.Font.GetYAdvance()
	maxLines := (maxY + int16(yAdv) - 1) / int16(yAdv)
	for l := len(lines); l > int(maxLines); l-- {
		lines = lines[1:]
	}
	w.sb.WriteString(strings.Join(lines, "\n"))

	liney := yAdv
	for _, s := range lines {
		w.writer.WriteLine(s, 0, int16(liney), draw.White)
		liney += yAdv
	}
	_ = w.display.Display()
}
