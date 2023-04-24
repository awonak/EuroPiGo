//go:build !pico && revision1
// +build !pico,revision1

package main

import (
	"log"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/internal/nonpico/events"
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
	clock.Init(module.Config{
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
	})

	events.SetupVoltageOutputListeners()
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
	err = europi.Bootstrap(
		europi.UsingEuroPi(europi.New(hal.Revision1)),
		europi.EnableDisplayLogger(false),
		europi.InitRandom(true),
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
		europi.UI(ui),
		europi.UIRefreshRate(time.Millisecond*50),
	)
	if err != nil {
		log.Fatalf("Bootstrap exited with: %v\n", err)
	}

	log.Println("done.")
}
