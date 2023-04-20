package main

import (
	"fmt"
	"math/rand"
	"time"

	clockgenerator "github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/internal/projects/randomskips/module"
	"github.com/heucuva/europi/units"
)

var (
	skip  module.RandomSkips
	clock clockgenerator.ClockGenerator

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

func startLoop() {
	setCV1 := panicCV("cv1")

	if err := skip.Init(module.Config{
		Gate: func(high bool) { // Gate 1
			if high {
				setCV1(1.0)
			} else {
				setCV1(0.0)
			}
		},
		Chance: 2.0 / 3.0,
	}); err != nil {
		panic(err)
	}

	if err := clock.Init(clockgenerator.Config{
		BPM:      120.0,
		Enabled:  true,
		ClockOut: skip.Gate,
	}); err != nil {
		panic(err)
	}
}

func mainLoop(deltaTime time.Duration) {
	clock.Tick(deltaTime)
	skip.Tick(deltaTime)
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
