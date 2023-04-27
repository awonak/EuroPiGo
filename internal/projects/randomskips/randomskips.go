package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	clockgenerator "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/module"
	clockScreen "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/screen"
	"github.com/awonak/EuroPiGo/internal/projects/randomskips/module"
	"github.com/awonak/EuroPiGo/internal/projects/randomskips/screen"
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

func makeGate(out hal.VoltageOutput) func(value bool) {
	return func(value bool) {
		if value {
			out.SetCV(1.0)
		} else {
			out.SetCV(0.0)
		}
	}
}

func appStart(e *europi.EuroPi) {
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

	e.DI.HandlerEx(hal.ChangeAny, func(value bool, _ time.Duration) {
		skip.Gate(value)
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
	if err := europi.Bootstrap(
		europi.EnableDisplayLogger(false),
		europi.InitRandom(true),
		europi.AppStart(appStart),
		europi.AppMainLoop(mainLoop),
		europi.AppMainLoopInterval(time.Millisecond*1),
		europi.UI(ui),
		europi.UIRefreshRate(time.Millisecond*50),
	); err != nil {
		panic(err)
	}
}
