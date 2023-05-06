package bootstrap

import (
	"errors"
	"time"

	europi "github.com/awonak/EuroPiGo"
)

// appHardwareWrapper sets up a wrapper around an app that expects a particular hardware interface
// this is for automated parameter interpretation
func appHardwareWrapper[THardware europi.Hardware](app any) any {
	start, _ := app.(ApplicationStart[THardware])
	mainLoop, _ := app.(ApplicationMainLoop[THardware])
	end, _ := app.(ApplicationEnd[THardware])
	return &appWrapper[THardware]{
		start:    start,
		mainLoop: mainLoop,
		end:      end,
	}
}

func getWrappedAppFuncs[THardware europi.Hardware](app any) (start AppStartFunc, mainLoop AppMainLoopFunc, end AppEndFunc) {
	appWrapper := appHardwareWrapper[THardware](app)
	if getStart, _ := appWrapper.(applicationStartProvider); getStart != nil {
		start = getStart.ApplicationStart()
	}

	if getMainLoop, _ := appWrapper.(applicationMainLoopProvider); getMainLoop != nil {
		mainLoop = getMainLoop.ApplicationMainLoop()
	}

	if getEnd, _ := appWrapper.(applicationEndProvider); getEnd != nil {
		end = getEnd.ApplicationEnd()
	}
	return
}

type applicationStartProvider interface {
	ApplicationStart() AppStartFunc
}

type applicationMainLoopProvider interface {
	ApplicationMainLoop() AppMainLoopFunc
}

type applicationEndProvider interface {
	ApplicationEnd() AppEndFunc
}

type appWrapper[THardware europi.Hardware] struct {
	start    ApplicationStart[THardware]
	mainLoop ApplicationMainLoop[THardware]
	end      ApplicationEnd[THardware]
}

func (a *appWrapper[THardware]) ApplicationStart() AppStartFunc {
	if a.start == nil {
		return nil
	}
	return a.doStart
}

func (a *appWrapper[THardware]) doStart(e europi.Hardware) {
	pi, ok := e.(THardware)
	if !ok {
		panic(errors.New("incorrect hardware type conversion"))
	}
	a.start.Start(pi)
}

func (a *appWrapper[THardware]) ApplicationMainLoop() AppMainLoopFunc {
	if a.mainLoop == nil {
		return nil
	}
	return a.doMainLoop
}

func (a *appWrapper[THardware]) doMainLoop(e europi.Hardware, deltaTime time.Duration) {
	pi, ok := e.(THardware)
	if !ok {
		panic(errors.New("incorrect hardware type conversion"))
	}
	a.mainLoop.MainLoop(pi, deltaTime)
}

func (a *appWrapper[THardware]) ApplicationEnd() AppEndFunc {
	if a.end == nil {
		return nil
	}
	return a.doEnd
}

func (a *appWrapper[THardware]) doEnd(e europi.Hardware) {
	pi, ok := e.(THardware)
	if !ok {
		panic(errors.New("incorrect hardware type conversion"))
	}
	a.end.End(pi)
}
