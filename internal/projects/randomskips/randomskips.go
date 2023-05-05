package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/bootstrap"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	clockgenerator "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/module"
	clockScreen "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/screen"
	"github.com/awonak/EuroPiGo/internal/projects/randomskips/module"
	"github.com/awonak/EuroPiGo/internal/projects/randomskips/screen"
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

type application struct {
	skip  *module.RandomSkips
	clock *clockgenerator.ClockGenerator

	ui             *screenbank.ScreenBank
	screenMain     screen.Main
	screenClock    clockScreen.Settings
	screenSettings screen.Settings
}

func newApplication() (*application, error) {
	skip := &module.RandomSkips{}
	clock := &clockgenerator.ClockGenerator{}

	app := &application{
		skip:  skip,
		clock: clock,
		screenMain: screen.Main{
			RandomSkips: skip,
			Clock:       clock,
		},
		screenClock: clockScreen.Settings{
			Clock: clock,
		},
		screenSettings: screen.Settings{
			RandomSkips: skip,
		},
	}

	var err error
	app.ui, err = screenbank.NewScreenBank(
		screenbank.WithScreen("main", "\u2b50", &app.screenMain),
		screenbank.WithScreen("settings", "\u2611", &app.screenSettings),
		screenbank.WithScreen("clock", "\u23f0", &app.screenClock),
	)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *application) Start(e *europi.EuroPi) {
	if err := app.skip.Init(module.Config{
		Gate:   makeGate(e.CV1),
		Chance: 0.5,
	}); err != nil {
		panic(err)
	}

	if err := app.clock.Init(clockgenerator.Config{
		BPM:      120.0,
		Enabled:  false,
		ClockOut: app.skip.Gate,
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(hal.ChangeAny, func(value bool, _ time.Duration) {
		app.skip.Gate(value)
	})
}

func (app *application) MainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	app.clock.Tick(deltaTime)
	app.skip.Tick(deltaTime)
}

func main() {
	app, err := newApplication()
	if err != nil {
		panic(err)
	}

	pi := europi.New()

	// some options shown below are being explicitly set to their defaults
	// only to showcase their existence.
	if err := bootstrap.Bootstrap(
		pi,
		bootstrap.EnableDisplayLogger(false),
		bootstrap.InitRandom(true),
		bootstrap.App(
			app,
			bootstrap.AppMainLoopInterval(time.Millisecond*1),
		),
		bootstrap.UI(
			app.ui,
			bootstrap.UIRefreshRate(time.Millisecond*50),
		),
	); err != nil {
		panic(err)
	}
}
