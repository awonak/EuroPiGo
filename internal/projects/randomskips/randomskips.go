package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	clockgenerator "github.com/heucuva/europi/internal/projects/clockgenerator/module"
	clockScreen "github.com/heucuva/europi/internal/projects/clockgenerator/screen"
	"github.com/heucuva/europi/internal/projects/randomskips/module"
	"github.com/heucuva/europi/internal/projects/randomskips/screen"
	"github.com/heucuva/europi/output"
)

var (
	skip  module.RandomSkips
	clock clockgenerator.ClockGenerator

	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		RandomSkips: &skip,
		Clock:       &clock,
	}
	screenClock = clockScreen.Settings{
		Clock: &clock,
	}
	screenSettings = screen.Settings{
		RandomSkips: &skip,
	}
)

func makeGate(out output.Output) func(high bool) {
	return func(high bool) {
		if high {
			out.On()
		} else {
			out.Off()
		}
	}
}

func startLoop(e *europi.EuroPi) {
	if err := skip.Init(module.Config{
		Gate:   makeGate(e.CV1),
		Chance: 0.5,
	}); err != nil {
		panic(err)
	}

	if err := clock.Init(clockgenerator.Config{
		BPM:      120.0,
		Enabled:  false,
		ClockOut: skip.Gate,
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.DI.Value()
		skip.Gate(high)
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	clock.Tick(deltaTime)
	skip.Tick(deltaTime)
}

func main() {
	var err error
	ui, err = screenbank.NewScreenBank(
		screenbank.WithScreen("main", "\u2b50", &screenMain),
		screenbank.WithScreen("settings", "\u2611", &screenSettings),
		screenbank.WithScreen("clock", "\u23f0", &screenClock),
	)
	if err != nil {
		panic(err)
	}

	// some options shown below are being explicitly set to their defaults
	// only to showcase their existence.
	europi.Bootstrap(
		europi.EnableDisplayLogger(false),
		europi.InitRandom(true),
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
		europi.UI(ui),
		europi.UIRefreshRate(time.Millisecond*50),
	)
}
