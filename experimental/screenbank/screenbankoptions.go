package screenbank

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
)

type ScreenBankOption func(sb *ScreenBank) error

// WithScreen sets up a new screen in the chain
//  logo is the emoji to use (see https://github.com/tinygo-org/tinyfont/blob/release/notoemoji/NotoEmoji-Regular-12pt.go)
func WithScreen(name string, logo string, screen europi.UserInterface) ScreenBankOption {
	return func(sb *ScreenBank) error {
		e := screenBankEntry{
			name:       name,
			logo:       logo,
			screen:     screen,
			enabled:    true,
			locked:     true,
			lastUpdate: time.Now(),
		}

		sb.bank = append(sb.bank, e)
		return nil
	}
}
