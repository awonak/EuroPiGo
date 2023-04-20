package europi

import (
	"machine"
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
	Button1(e *EuroPi, p machine.Pin)
}

type UserInterfaceButton1Debounce interface {
	Button1Debounce() time.Duration
}

type UserInterfaceButton1Ex interface {
	Button1Ex(e *EuroPi, p machine.Pin, high bool)
}

type UserInterfaceButton1Long interface {
	Button1Long(e *EuroPi, p machine.Pin)
}

type UserInterfaceButton2 interface {
	Button2(e *EuroPi, p machine.Pin)
}

type UserInterfaceButton2Debounce interface {
	Button2Debounce() time.Duration
}

type UserInterfaceButton2Ex interface {
	Button2Ex(e *EuroPi, p machine.Pin, high bool)
}

type UserInterfaceButton2Long interface {
	Button2Long(e *EuroPi, p machine.Pin)
}

var (
	ui uiModule
)

func enableUI(e *EuroPi, screen UserInterface, interval time.Duration) {
	ui.screen = screen
	if ui.screen == nil {
		return
	}

	ui.logoPainter, _ = screen.(UserInterfaceLogoPainter)

	ui.repaint = make(chan struct{}, 1)

	var (
		inputB1  func(e *EuroPi, p machine.Pin, high bool)
		inputB1L func(e *EuroPi, p machine.Pin)
	)
	if in, ok := screen.(UserInterfaceButton1); ok {
		var debounceDelay time.Duration
		if db, ok := screen.(UserInterfaceButton1Debounce); ok {
			debounceDelay = db.Button1Debounce()
		}
		var lastTrigger time.Time
		inputB1 = func(e *EuroPi, p machine.Pin, high bool) {
			now := time.Now()
			if !high && (debounceDelay == 0 || now.Sub(lastTrigger) >= debounceDelay) {
				lastTrigger = now
				in.Button1(e, p)
			}
		}
	} else if in, ok := screen.(UserInterfaceButton1Ex); ok {
		inputB1 = in.Button1Ex
	}
	if in, ok := screen.(UserInterfaceButton1Long); ok {
		inputB1L = in.Button1Long
	}
	ui.setupButton(e, e.B1, inputB1, inputB1L)

	var (
		inputB2  func(e *EuroPi, p machine.Pin, high bool)
		inputB2L func(e *EuroPi, p machine.Pin)
	)
	if in, ok := screen.(UserInterfaceButton2); ok {
		var debounceDelay time.Duration
		if db, ok := screen.(UserInterfaceButton2Debounce); ok {
			debounceDelay = db.Button2Debounce()
		}
		var lastTrigger time.Time
		inputB2 = func(e *EuroPi, p machine.Pin, high bool) {
			now := time.Now()
			if !high && (debounceDelay == 0 || now.Sub(lastTrigger) >= debounceDelay) {
				lastTrigger = now
				in.Button2(e, p)
			}
		}
	} else if in, ok := screen.(UserInterfaceButton2Ex); ok {
		inputB2 = in.Button2Ex
	}
	if in, ok := screen.(UserInterfaceButton2Long); ok {
		inputB2L = in.Button2Long
	}
	ui.setupButton(e, e.B2, inputB2, inputB2L)

	ui.wg.Add(1)
	go ui.run(e, interval)
}

func startUI(e *EuroPi) {
	if ui.screen == nil {
		return
	}

	ui.screen.Start(e)
}

// ForceRepaintUI schedules a forced repaint of the UI (if it is configured and running)
func ForceRepaintUI(e *EuroPi) {
	if ui.repaint != nil {
		ui.repaint <- struct{}{}
	}
}

func disableUI(e *EuroPi) {
	if ui.stop != nil {
		ui.stop()
	}

	if ui.repaint != nil {
		close(ui.repaint)
	}

	ui.wait()
}
