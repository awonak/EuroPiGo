package screen

import (
	"fmt"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	clockgenerator "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/module"
	"github.com/awonak/EuroPiGo/internal/projects/randomskips/module"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont/proggy"
)

type Main struct {
	RandomSkips *module.RandomSkips
	Clock       *clockgenerator.ClockGenerator
	writer      fontwriter.Writer
}

const (
	line1y int16 = 11
	line2y int16 = 23
)

var (
	DefaultFont = &proggy.TinySZ8pt7b
)

func (m *Main) Start(e *europi.EuroPi) {
	m.writer = fontwriter.Writer{
		Display: e.OLED,
		Font:    DefaultFont,
	}
}

func (m *Main) Button1Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Main) Button1(e *europi.EuroPi, deltaTime time.Duration) {
	m.Clock.Toggle()
}

func (m *Main) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	if m.Clock.Enabled() {
		tinydraw.Line(m.writer.Display, 0, 0, 7, 0, draw.White)
	}
	m.writer.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f 3:%2.1f", e.CV1.Voltage(), e.CV2.Voltage(), e.CV3.Voltage()), 0, line1y, draw.White)
	m.writer.WriteLine(fmt.Sprintf("4:%2.1f 5:%2.1f 6:%2.1f", e.CV4.Voltage(), e.CV5.Voltage(), e.CV6.Voltage()), 0, line2y, draw.White)
}
