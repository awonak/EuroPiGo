package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/internal/projects/clockgenerator/module"
	"github.com/awonak/EuroPiGo/internal/projects/clockgenerator/screen"
)

var (
	clock      module.ClockGenerator
	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		Clock: &clock,
	}
	screenSettings = screen.Settings{
		Clock: &clock,
	}
)

func startLoop(e *europi.EuroPi) {
	if err := clock.Init(module.Config{
		BPM:          120.0,
		GateDuration: time.Millisecond * 100,
		Enabled:      true,
		ClockOut: func(value bool) {
			if value {
				e.CV1.SetCV(1.0)
			} else {
				e.CV1.SetCV(0.0)
			}
			europi.ForceRepaintUI(e)
		},
	}); err != nil {
		panic(err)
	}
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	clock.Tick(deltaTime)
}

func main() {
	var err error
	ui, err = screenbank.NewScreenBank(
		screenbank.WithScreen("main", "\u2b50", &screenMain),
		screenbank.WithScreen("settings", "\u2611", &screenSettings),
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
