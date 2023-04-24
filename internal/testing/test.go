//go:build !pico && test
// +build !pico,test

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
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

func panicHandler(e *europi.EuroPi, reason any) {
	log.Fatalln(reason)
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

	// set up bus subscriptions
	bus := rev1.DefaultEventBus
	go func() {
		diValueTicker := time.NewTicker(time.Second * 5)
		defer diValueTicker.Stop()
		diState := false

		aiValueTicker := time.NewTicker(time.Second * 4)
		defer aiValueTicker.Stop()

		b1Ticker := time.NewTicker(time.Second * 1)
		defer b1Ticker.Stop()
		b1State := false

		b2Ticker := time.NewTicker(time.Second * 3)
		defer b2Ticker.Stop()
		b2State := false

		for {
			select {
			case <-diValueTicker.C:
				value := rand.Float32() < 0.5
				bus.Post(fmt.Sprintf("hw_value_%d", hal.HardwareIdDigital1Input), rev1.HwMessageDigitalValue{
					Value: value,
				})

				if diState != value {
					diState = value
					if value {
						// rising
						bus.Post(fmt.Sprintf("hw_interrupt_%d", hal.HardwareIdDigital1Input), rev1.HwMessageInterrupt{
							Change: hal.ChangeRising,
						})
					} else {
						// falling
						bus.Post(fmt.Sprintf("hw_interrupt_%d", hal.HardwareIdDigital1Input), rev1.HwMessageInterrupt{
							Change: hal.ChangeFalling,
						})
					}
				}

			case <-aiValueTicker.C:
				bus.Post(fmt.Sprintf("hw_value_%d", hal.HardwareIdAnalog1Input), rev1.HwMessageADCValue{
					Value: uint16(rand.Int31n(math.MaxUint16)),
				})

			case <-b1Ticker.C:
				value := rand.Float32() < 0.5
				bus.Post(fmt.Sprintf("hw_value_%d", hal.HardwareIdButton1Input), rev1.HwMessageDigitalValue{
					Value: value,
				})

				if b1State != value {
					b1State = value
					if value {
						// rising
						bus.Post(fmt.Sprintf("hw_interrupt_%d", hal.HardwareIdButton1Input), rev1.HwMessageInterrupt{
							Change: hal.ChangeRising,
						})
					} else {
						// falling
						bus.Post(fmt.Sprintf("hw_interrupt_%d", hal.HardwareIdButton1Input), rev1.HwMessageInterrupt{
							Change: hal.ChangeFalling,
						})
					}
				}

			case <-b2Ticker.C:
				value := rand.Float32() < 0.5
				bus.Post(fmt.Sprintf("hw_value_%d", hal.HardwareIdButton2Input), rev1.HwMessageDigitalValue{
					Value: value,
				})

				if b2State != value {
					b2State = value
					if value {
						// rising
						bus.Post(fmt.Sprintf("hw_interrupt_%d", hal.HardwareIdButton2Input), rev1.HwMessageInterrupt{
							Change: hal.ChangeRising,
						})
					} else {
						// falling
						bus.Post(fmt.Sprintf("hw_interrupt_%d", hal.HardwareIdButton2Input), rev1.HwMessageInterrupt{
							Change: hal.ChangeFalling,
						})
					}
				}
			}
		}
	}()

	for id := hal.HardwareIdVoltage1Output; id <= hal.HardwareIdVoltage6Output; id++ {
		fn := func(hid hal.HardwareId) func(rev1.HwMessagePwmValue) {
			return func(msg rev1.HwMessagePwmValue) {
				log.Printf("CV%d: %v", hid-hal.HardwareIdVoltage1Output+1, msg.Value)
			}
		}(id)
		event.Subscribe(bus, fmt.Sprintf("hw_pwm_%d", id), fn)
	}

	event.Subscribe(bus, fmt.Sprintf("hw_display_%d", hal.HardwareIdDisplay1Output), func(msg rev1.HwMessageDisplay) {
		if msg.Op == 1 {
			return
		}
		log.Printf("display: %v(%+v)", msg.Op, msg.Operands)
	})

	// some options shown below are being explicitly set to their defaults
	// only to showcase their existence.
	europi.Bootstrap(
		europi.EnableDisplayLogger(false),
		//europi.BeginDestroy(panicHandler),
		europi.InitRandom(true),
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
		europi.UI(ui),
		europi.UIRefreshRate(time.Millisecond*50),
	)
}
