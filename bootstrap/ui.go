package bootstrap

import (
	"context"
	"time"

	europi "github.com/awonak/EuroPiGo"
)

type UserInterface[THardware europi.Hardware] interface {
	Start(e THardware)
	Paint(e THardware, deltaTime time.Duration)
}

type UserInterfaceLogoPainter[THardware europi.Hardware] interface {
	PaintLogo(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton1[THardware europi.Hardware] interface {
	Button1(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton1Debounce interface {
	Button1Debounce() time.Duration
}

type UserInterfaceButton1Ex[THardware europi.Hardware] interface {
	Button1Ex(e THardware, value bool, deltaTime time.Duration)
}

type UserInterfaceButton1Long[THardware europi.Hardware] interface {
	Button1Long(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton2[THardware europi.Hardware] interface {
	Button2(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton2Debounce interface {
	Button2Debounce() time.Duration
}

type UserInterfaceButton2Ex[THardware europi.Hardware] interface {
	Button2Ex(e THardware, value bool, deltaTime time.Duration)
}

type UserInterfaceButton2Long[THardware europi.Hardware] interface {
	Button2Long(e THardware, deltaTime time.Duration)
}

var (
	ui uiModule
)

func enableUI(ctx context.Context, e europi.Hardware, config bootstrapUIConfig) {
	ui.setup(e, config.ui)

	ui.start(ctx, e, config.uiRefreshRate)
}

func startUI(e europi.Hardware) {
	if ui.screen == nil {
		return
	}

	ui.screen.Start(e)
}

// ForceRepaintUI schedules a forced repaint of the UI (if it is configured and running)
func ForceRepaintUI(e europi.Hardware) {
	ui.repaint()
}

func disableUI(e europi.Hardware) {
	ui.shutdown()
}
