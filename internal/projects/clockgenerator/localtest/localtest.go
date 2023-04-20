package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/units"
)

var (
	clockgenerator module.ClockGenerator

	outMap map[string]float32 = make(map[string]float32)
)

func reportVolts(name string, v float32) {
	switch name {
	//case "cv1":
	case "cv2":
	case "cv3":
	case "cv4":
	case "cv5":
	case "cv6":
	default:
		old := outMap[name]
		if old != v {
			fmt.Printf("%s: %v Volts\n", name, v)
			outMap[name] = v
		}
	}
}

func panicVOct(name string) func(units.VOct) {
	return func(voct units.VOct) {
		v := voct.ToVolts()
		reportVolts(name, v)
	}
}

func panicCV(name string) func(units.CV) {
	return func(cv units.CV) {
		v := cv.ToVolts()
		reportVolts(name, v)
	}
}

func bipolarOut(out func(units.CV)) func(cv units.BipolarCV) {
	return func(cv units.BipolarCV) {
		out(cv.ToCV())
	}
}

func startLoop() {
	setCV1 := panicCV("cv1")

	if err := clockgenerator.Init(module.Config{
		ClockOut: func(high bool) {
			if high {
				setCV1(1)
			} else {
				setCV1(0)
			}
		},
		BPM:          120.0,
		GateDuration: time.Millisecond * 100,
		Enabled:      true,
	}); err != nil {
		panic(err)
	}
}

func mainLoop(deltaTime time.Duration) {
	clockgenerator.Tick(deltaTime)
}

func main() {
	startLoop()

	ticker := time.NewTicker(time.Millisecond * 1)
	defer ticker.Stop()
	prev := time.Now()
	for {
		now := <-ticker.C
		mainLoop(now.Sub(prev))
		prev = now
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
