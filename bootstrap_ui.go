package europi

import (
	"context"
	"time"
)

type UserInterface[THardware Hardware] interface {
	Start(e THardware)
	Paint(e THardware, deltaTime time.Duration)
}

type UserInterfaceLogoPainter[THardware Hardware] interface {
	PaintLogo(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton1[THardware Hardware] interface {
	Button1(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton1Debounce interface {
	Button1Debounce() time.Duration
}

type UserInterfaceButton1Ex[THardware Hardware] interface {
	Button1Ex(e THardware, value bool, deltaTime time.Duration)
}

type UserInterfaceButton1Long[THardware Hardware] interface {
	Button1Long(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton2[THardware Hardware] interface {
	Button2(e THardware, deltaTime time.Duration)
}

type UserInterfaceButton2Debounce interface {
	Button2Debounce() time.Duration
}

type UserInterfaceButton2Ex[THardware Hardware] interface {
	Button2Ex(e THardware, value bool, deltaTime time.Duration)
}

type UserInterfaceButton2Long[THardware Hardware] interface {
	Button2Long(e THardware, deltaTime time.Duration)
}

var (
	ui uiModule
)

func enableUI(ctx context.Context, e Hardware, config bootstrapUIConfig) {
	ui.setup(e, config.ui)

	ui.start(ctx, e, config.uiRefreshRate)
}

func startUI(e Hardware) {
	if ui.screen == nil {
		return
	}

	ui.screen.Start(e)
}

// ForceRepaintUI schedules a forced repaint of the UI (if it is configured and running)
func ForceRepaintUI(e Hardware) {
	ui.repaint()
}

func disableUI(e Hardware) {
	ui.shutdown()
}
