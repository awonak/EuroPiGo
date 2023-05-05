package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	"tinygo.org/x/tinyfont/notoemoji"
)

type ScreenBank struct {
	current int
	bank    []entry
	writer  fontwriter.Writer
}

var (
	DefaultFont = &notoemoji.NotoEmojiRegular12pt
)

func NewScreenBank(opts ...ScreenBankOption) (*ScreenBank, error) {
	sb := &ScreenBank{
		writer: fontwriter.Writer{
			Font: DefaultFont,
		},
	}

	for _, opt := range opts {
		if err := opt(sb); err != nil {
			return nil, err
		}
	}

	return sb, nil
}

func (sb *ScreenBank) CurrentName() string {
	if len(sb.bank) == 0 {
		return ""
	}
	return sb.bank[sb.current].name
}

func (sb *ScreenBank) Current() *entryWrapper[europi.Hardware] {
	if len(sb.bank) == 0 {
		return nil
	}
	return &sb.bank[sb.current].screen
}

func (sb *ScreenBank) transitionTo(idx int) {
	if sb.current >= len(sb.bank) || len(sb.bank) == 0 {
		return
	}

	cur := sb.bank[sb.current]
	cur.lock()
	sb.current = idx
	if sb.current >= len(sb.bank) {
		sb.current = 0
	}
	sb.bank[sb.current].unlock()
}

func (sb *ScreenBank) Goto(idx int) {
	sb.transitionTo(idx)
}

func (sb *ScreenBank) GotoNamed(name string) {
	for i, screen := range sb.bank {
		if screen.name == name {
			sb.transitionTo(i)
			return
		}
	}
}

func (sb *ScreenBank) Next() {
	sb.transitionTo(sb.current + 1)
}

func (sb *ScreenBank) Start(e europi.Hardware) {
	for i := range sb.bank {
		s := &sb.bank[i]

		s.lock()
		s.screen.screen.Start(e)
		s.lastUpdate = time.Now()
		s.unlock()
	}
}

func (sb *ScreenBank) PaintLogo(e europi.Hardware, deltaTime time.Duration) {
	display := europi.Display(e)
	if sb.current >= len(sb.bank) || display == nil {
		return
	}

	cur := &sb.bank[sb.current]
	cur.lock()
	if cur.logo != "" {
		sb.writer.Display = display
		sb.writer.WriteLineInverseAligned(cur.logo, 0, 16, draw.White, fontwriter.AlignRight, fontwriter.AlignMiddle)
	}
	cur.unlock()
}

func (sb *ScreenBank) Paint(e europi.Hardware, deltaTime time.Duration) {
	if sb.current >= len(sb.bank) {
		return
	}

	cur := &sb.bank[sb.current]
	cur.lock()
	now := time.Now()
	cur.screen.screen.Paint(e, now.Sub(cur.lastUpdate))
	cur.lastUpdate = now
	cur.unlock()
}

func (sb *ScreenBank) Button1Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	screen := sb.Current()
	if cur := screen.button1; cur != nil {
		if !value {
			cur.Button1(e, deltaTime)
		}
	} else if cur := screen.button1Ex; cur != nil {
		cur.Button1Ex(e, value, deltaTime)
	}
}

func (sb *ScreenBank) Button1Long(e europi.Hardware, deltaTime time.Duration) {
	screen := sb.Current()
	if cur := screen.button1Long; cur != nil {
		cur.Button1Long(e, deltaTime)
	} else {
		// try the short-press
		sb.Button1Ex(e, false, deltaTime)
	}
}

func (sb *ScreenBank) Button2Ex(e europi.Hardware, value bool, deltaTime time.Duration) {
	screen := sb.Current()
	if cur := screen.button2; cur != nil {
		if !value {
			cur.Button2(e, deltaTime)
		}
	} else if cur := screen.button2Ex; cur != nil {
		cur.Button2Ex(e, value, deltaTime)
	}
}

func (sb *ScreenBank) Button2Long(e europi.Hardware, deltaTime time.Duration) {
	sb.Next()
}
