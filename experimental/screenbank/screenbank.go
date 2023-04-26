package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	"tinygo.org/x/tinyfont/notoemoji"
)

type ScreenBank struct {
	screen  europi.UserInterface
	current int
	bank    []screenBankEntry
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

func (sb *ScreenBank) Current() europi.UserInterface {
	if len(sb.bank) == 0 {
		return nil
	}
	return sb.bank[sb.current].screen
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

func (sb *ScreenBank) Start(e *europi.EuroPi) {
	for i := range sb.bank {
		s := &sb.bank[i]

		s.lock()
		s.screen.Start(e)
		s.lastUpdate = time.Now()
		s.unlock()
	}
}

func (sb *ScreenBank) PaintLogo(e *europi.EuroPi, deltaTime time.Duration) {
	if sb.current >= len(sb.bank) {
		return
	}

	cur := &sb.bank[sb.current]
	cur.lock()
	if cur.logo != "" {
		sb.writer.Display = e.Display
		sb.writer.WriteLineInverseAligned(cur.logo, 0, 16, draw.White, fontwriter.AlignRight, fontwriter.AlignMiddle)
	}
	cur.unlock()
}

func (sb *ScreenBank) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	if sb.current >= len(sb.bank) {
		return
	}

	cur := &sb.bank[sb.current]
	cur.lock()
	now := time.Now()
	cur.screen.Paint(e, now.Sub(cur.lastUpdate))
	cur.lastUpdate = now
	cur.unlock()
}

func (sb *ScreenBank) Button1Ex(e *europi.EuroPi, value bool, deltaTime time.Duration) {
	screen := sb.Current()
	if cur, ok := screen.(europi.UserInterfaceButton1); ok {
		if !value {
			cur.Button1(e, deltaTime)
		}
	} else if cur, ok := screen.(europi.UserInterfaceButton1Ex); ok {
		cur.Button1Ex(e, value, deltaTime)
	}
}

func (sb *ScreenBank) Button1Long(e *europi.EuroPi, deltaTime time.Duration) {
	if cur, ok := sb.Current().(europi.UserInterfaceButton1Long); ok {
		cur.Button1Long(e, deltaTime)
	} else {
		// try the short-press
		sb.Button1Ex(e, false, deltaTime)
	}
}

func (sb *ScreenBank) Button2Ex(e *europi.EuroPi, value bool, deltaTime time.Duration) {
	screen := sb.Current()
	if cur, ok := screen.(europi.UserInterfaceButton2); ok {
		if !value {
			cur.Button2(e, deltaTime)
		}
	} else if cur, ok := screen.(europi.UserInterfaceButton2Ex); ok {
		cur.Button2Ex(e, value, deltaTime)
	}
}

func (sb *ScreenBank) Button2Long(e *europi.EuroPi, deltaTime time.Duration) {
	sb.Next()
}
