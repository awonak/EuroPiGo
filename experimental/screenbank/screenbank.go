package screenbank

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/output"
)

type ScreenBank struct {
	screen  europi.UserInterface
	current int
	bank    []screenBankEntry
}

func NewScreenBank(opts ...ScreenBankOption) (*ScreenBank, error) {
	sb := &ScreenBank{}

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
		e.Display.WriteEmojiLineInverseAligned(cur.logo, 0, 16, output.AlignRight, output.AlignMiddle)
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

func (sb *ScreenBank) Button1Ex(e *europi.EuroPi, p machine.Pin, high bool) {
	screen := sb.Current()
	if cur, ok := screen.(europi.UserInterfaceButton1); ok {
		if !high {
			cur.Button1(e, p)
		}
	} else if cur, ok := screen.(europi.UserInterfaceButton1Ex); ok {
		cur.Button1Ex(e, p, high)
	}
}

func (sb *ScreenBank) Button1Long(e *europi.EuroPi, p machine.Pin) {
	if cur, ok := sb.Current().(europi.UserInterfaceButton1Long); ok {
		cur.Button1Long(e, p)
	} else {
		// try the short-press
		sb.Button1Ex(e, p, false)
	}
}

func (sb *ScreenBank) Button2Ex(e *europi.EuroPi, p machine.Pin, high bool) {
	screen := sb.Current()
	if cur, ok := screen.(europi.UserInterfaceButton2); ok {
		if !high {
			cur.Button2(e, p)
		}
	} else if cur, ok := screen.(europi.UserInterfaceButton2Ex); ok {
		cur.Button2Ex(e, p, high)
	}
}

func (sb *ScreenBank) Button2Long(e *europi.EuroPi, p machine.Pin) {
	sb.Next()
}
