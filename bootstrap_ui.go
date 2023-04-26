package europi

import (
	"time"
)

type UserInterface interface {
	Start(e *EuroPi)
	Paint(e *EuroPi, deltaTime time.Duration)
}

type UserInterfaceLogoPainter interface {
	PaintLogo(e *EuroPi, deltaTime time.Duration)
}

type UserInterfaceButton1 interface {
	Button1(e *EuroPi, deltaTime time.Duration)
}

type UserInterfaceButton1Debounce interface {
	Button1Debounce() time.Duration
}

type UserInterfaceButton1Ex interface {
	Button1Ex(e *EuroPi, value bool, deltaTime time.Duration)
}

type UserInterfaceButton1Long interface {
	Button1Long(e *EuroPi, deltaTime time.Duration)
}

type UserInterfaceButton2 interface {
	Button2(e *EuroPi, deltaTime time.Duration)
}

type UserInterfaceButton2Debounce interface {
	Button2Debounce() time.Duration
}

type UserInterfaceButton2Ex interface {
	Button2Ex(e *EuroPi, value bool, deltaTime time.Duration)
}

type UserInterfaceButton2Long interface {
	Button2Long(e *EuroPi, deltaTime time.Duration)
}

var (
	ui uiModule
)

func enableUI(e *EuroPi, screen UserInterface, interval time.Duration) {
	ui.setup(e, screen)

	ui.start(e, interval)
}

func startUI(e *EuroPi) {
	if ui.screen == nil {
		return
	}

	ui.screen.Start(e)
}

// ForceRepaintUI schedules a forced repaint of the UI (if it is configured and running)
func ForceRepaintUI(e *EuroPi) {
	ui.repaint()
}

func disableUI(e *EuroPi) {
	ui.shutdown()
}
