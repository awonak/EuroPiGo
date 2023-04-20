package displaylogger

import (
	"strings"

	"github.com/heucuva/europi/output"
)

type Logger struct {
	sb      strings.Builder
	Display *output.Display
}

func (w *Logger) repaint() {
	str := w.sb.String()

	fnt := output.DefaultFont
	w.Display.SetFont(fnt)
	w.Display.ClearBuffer()

	lines := strings.Split(str, "\n")
	w.sb.Reset()
	_, maxY := w.Display.Size()
	maxLines := (maxY + int16(fnt.YAdvance) - 1) / int16(fnt.YAdvance)
	for l := len(lines); l > int(maxLines); l-- {
		lines = lines[1:]
	}
	w.sb.WriteString(strings.Join(lines, "\n"))

	liney := fnt.YAdvance
	for _, s := range lines {
		w.Display.WriteLine(s, 0, int16(liney))
		liney += fnt.YAdvance
	}
	_ = w.Display.Display()
}

func (w *Logger) Write(p []byte) (n int, err error) {
	n, err = w.sb.Write(p)
	if err != nil {
		return
	}

	w.repaint()
	return
}

func (w *Logger) Flush() {
	w.repaint()
}
