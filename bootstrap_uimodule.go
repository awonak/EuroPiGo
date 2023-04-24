package europi

import (
	"context"
	"sync"
	"time"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

type uiModule struct {
	screen      UserInterface
	logoPainter UserInterfaceLogoPainter
	repaint     chan struct{}
	stop        context.CancelFunc
	wg          sync.WaitGroup
}

func (u *uiModule) wait() {
	u.wg.Wait()
}

func (u *uiModule) run(e *EuroPi, interval time.Duration) {
	defer u.wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	ui.stop = cancel
	defer ui.stop()

	t := time.NewTicker(interval)
	defer t.Stop()

	disp := e.Display
	lastTime := time.Now()

	paint := func(now time.Time) {
		deltaTime := now.Sub(lastTime)
		lastTime = now
		disp.ClearBuffer()
		if u.logoPainter != nil {
			u.logoPainter.PaintLogo(e, deltaTime)
		}
		u.screen.Paint(e, deltaTime)
		_ = disp.Display()
	}

	for {
		select {
		case <-ctx.Done():
			return

		case <-ui.repaint:
			paint(time.Now())

		case now := <-t.C:
			paint(now)
		}
	}
}

func (u *uiModule) setupButton(e *EuroPi, btn hal.ButtonInput, onShort func(e *EuroPi, value bool, deltaTime time.Duration), onLong func(e *EuroPi, deltaTime time.Duration)) {
	if btn == nil {
		return
	}

	if onShort == nil && onLong == nil {
		return
	}

	if onShort == nil {
		// no-op
		onShort = func(e *EuroPi, value bool, deltaTime time.Duration) {}
	}

	// if no long-press handler present, just reuse short-press handler
	if onLong == nil {
		onLong = func(e *EuroPi, deltaTime time.Duration) {
			onShort(e, false, deltaTime)
		}
	}

	const longDuration = time.Millisecond * 650

	btn.HandlerEx(hal.ChangeAny, func(value bool, deltaTime time.Duration) {
		if value {
			onShort(e, value, deltaTime)
		} else if deltaTime < longDuration {
			onShort(e, value, deltaTime)
		} else {
			onLong(e, deltaTime)
		}
	})
}
