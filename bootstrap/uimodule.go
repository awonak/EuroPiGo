package bootstrap

import (
	"context"
	"sync"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/debounce"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// LongPressDuration is the amount of time a button is in a held/pressed state before
// it is considered to be a 'long' press.
// TODO: This is eventually intended to be a persisted setting, configurable by the user.
const LongPressDuration = time.Millisecond * 650

type uiModule struct {
	screen      UserInterface[europi.Hardware]
	logoPainter UserInterfaceLogoPainter[europi.Hardware]
	repaintCh   chan struct{}
	stop        context.CancelFunc
	wg          sync.WaitGroup
}

func (u *uiModule) setup(e europi.Hardware, screen UserInterface[europi.Hardware]) {
	b1 := europi.Button(e, 0)
	b2 := europi.Button(e, 1)

	ui.screen = screen
	if ui.screen == nil {
		return
	}

	ui.logoPainter, _ = screen.(UserInterfaceLogoPainter[europi.Hardware])

	ui.repaintCh = make(chan struct{}, 1)

	var (
		inputB1  func(e europi.Hardware, value bool, deltaTime time.Duration)
		inputB1L func(e europi.Hardware, deltaTime time.Duration)
	)
	if in, ok := screen.(UserInterfaceButton1[europi.Hardware]); ok {
		var debounceDelay time.Duration
		if db, ok := screen.(UserInterfaceButton1Debounce); ok {
			debounceDelay = db.Button1Debounce()
		}
		inputDB := debounce.NewDebouncer(func(value bool, deltaTime time.Duration) {
			if !value {
				in.Button1(e, deltaTime)
			}
		}).Debounce(debounceDelay)
		inputB1 = func(e europi.Hardware, value bool, deltaTime time.Duration) {
			inputDB(value)
		}
	} else if in, ok := screen.(UserInterfaceButton1Ex[europi.Hardware]); ok {
		inputB1 = in.Button1Ex
	}
	if in, ok := screen.(UserInterfaceButton1Long[europi.Hardware]); ok {
		inputB1L = in.Button1Long
	}
	ui.setupButton(e, b1, inputB1, inputB1L)

	var (
		inputB2  func(e europi.Hardware, value bool, deltaTime time.Duration)
		inputB2L func(e europi.Hardware, deltaTime time.Duration)
	)
	if in, ok := screen.(UserInterfaceButton2[europi.Hardware]); ok {
		var debounceDelay time.Duration
		if db, ok := screen.(UserInterfaceButton2Debounce); ok {
			debounceDelay = db.Button2Debounce()
		}
		inputDB := debounce.NewDebouncer(func(value bool, deltaTime time.Duration) {
			if !value {
				in.Button2(e, deltaTime)
			}
		}).Debounce(debounceDelay)
		inputB2 = func(e europi.Hardware, value bool, deltaTime time.Duration) {
			inputDB(value)
		}
	} else if in, ok := screen.(UserInterfaceButton2Ex[europi.Hardware]); ok {
		inputB2 = in.Button2Ex
	}
	if in, ok := screen.(UserInterfaceButton2Long[europi.Hardware]); ok {
		inputB2L = in.Button2Long
	}
	ui.setupButton(e, b2, inputB2, inputB2L)
}

func (u *uiModule) start(ctx context.Context, e europi.Hardware, interval time.Duration) {
	ui.wg.Add(1)
	go ui.run(ctx, e, interval)
}

func (u *uiModule) wait() {
	u.wg.Wait()
}

func (u *uiModule) repaint() {
	if u.repaintCh != nil {
		u.repaintCh <- struct{}{}
	}
}

func (u *uiModule) shutdown() {
	if u.stop != nil {
		u.stop()
	}

	if ui.repaintCh != nil {
		close(ui.repaintCh)
	}

	ui.wait()
}

func (u *uiModule) run(ctx context.Context, e europi.Hardware, interval time.Duration) {
	defer u.wg.Done()

	disp := europi.Display(e)
	if disp == nil {
		// no display means no ui
		// TODO: make uiModule work when any user input/output is specified, not just display
		return
	}

	myCtx, cancel := context.WithCancel(ctx)
	ui.stop = cancel
	defer ui.stop()

	t := time.NewTicker(interval)
	defer t.Stop()

	paint := func(deltaTime time.Duration) {
		disp.ClearBuffer()
		if u.logoPainter != nil {
			u.logoPainter.PaintLogo(e, deltaTime)
		}
		u.screen.Paint(e, deltaTime)
		_ = disp.Display()
	}

	lastTime := time.Now()
	for {
		select {
		case <-myCtx.Done():
			return

		case <-ui.repaintCh:
			now := time.Now()
			deltaTime := now.Sub(lastTime)
			lastTime = now
			paint(deltaTime)

		case now := <-t.C:
			deltaTime := now.Sub(lastTime)
			lastTime = now
			paint(deltaTime)
		}
	}
}

func (u *uiModule) setupButton(e europi.Hardware, btn hal.ButtonInput, onShort func(e europi.Hardware, value bool, deltaTime time.Duration), onLong func(e europi.Hardware, deltaTime time.Duration)) {
	if btn == nil {
		return
	}

	if onShort == nil && onLong == nil {
		return
	}

	if onShort == nil {
		// no-op
		onShort = func(e europi.Hardware, value bool, deltaTime time.Duration) {}
	}

	// if no long-press handler present, just reuse short-press handler
	if onLong == nil {
		onLong = func(e europi.Hardware, deltaTime time.Duration) {
			onShort(e, false, deltaTime)
		}
	}

	btn.HandlerEx(hal.ChangeAny, func(value bool, deltaTime time.Duration) {
		if value {
			onShort(e, value, deltaTime)
		} else if deltaTime < LongPressDuration {
			onShort(e, value, deltaTime)
		} else {
			onLong(e, deltaTime)
		}
	})
}
