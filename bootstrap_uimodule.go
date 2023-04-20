package europi

import (
	"context"
	"machine"
	"sync"
	"time"

	"github.com/heucuva/europi/input"
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
		disp.Display()
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

func (u *uiModule) setupButton(e *EuroPi, btn input.DigitalReader, onShort func(e *EuroPi, p machine.Pin, high bool), onLong func(e *EuroPi, p machine.Pin)) {
	if onShort == nil && onLong == nil {
		return
	}

	if onShort == nil {
		// no-op
		onShort = func(e *EuroPi, p machine.Pin, high bool) {}
	}

	// if no long-press handler present, just reuse short-press handler
	if onLong == nil {
		onLong = func(e *EuroPi, p machine.Pin) {
			onShort(e, p, false)
		}
	}

	const longDuration = time.Millisecond * 650

	btn.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := btn.Value()
		if high {
			onShort(e, p, high)
		} else {
			startDown := btn.LastChange()
			deltaTime := time.Since(startDown)
			if deltaTime < longDuration {
				onShort(e, p, high)
			} else {
				onLong(e, p)
			}
		}
	})
}
