package europi

import (
	"time"
)

// appHardwareWrapper sets up a wrapper around an app that expects a particular hardware interface
// this is for automated parameter interpretation
func appHardwareWrapper[THardware Hardware](app any) any {
	start, _ := app.(ApplicationStart[THardware])
	mainLoop, _ := app.(ApplicationMainLoop[THardware])
	end, _ := app.(ApplicationEnd[THardware])
	return &appWrapper[THardware]{
		start:    start,
		mainLoop: mainLoop,
		end:      end,
	}
}

func getAppFuncs(app any) (start AppStartFunc, mainLoop AppMainLoopFunc, end AppEndFunc) {
	if appStart, _ := app.(ApplicationStart[Hardware]); appStart != nil {
		start = appStart.Start
	}
	if appMainLoop, _ := app.(ApplicationMainLoop[Hardware]); appMainLoop != nil {
		mainLoop = appMainLoop.MainLoop
	}
	if appEnd, _ := app.(ApplicationEnd[Hardware]); appEnd != nil {
		end = appEnd.End
	}

	if start == nil && mainLoop == nil && end == nil {
		start, mainLoop, end = getWrappedAppFuncs[*EuroPiPrototype](app)
	}

	if start == nil && mainLoop == nil && end == nil {
		start, mainLoop, end = getWrappedAppFuncs[*EuroPi](app)
	}
	return
}

func getWrappedAppFuncs[THardware Hardware](app any) (start AppStartFunc, mainLoop AppMainLoopFunc, end AppEndFunc) {
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

type appWrapper[THardware Hardware] struct {
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

func (a *appWrapper[THardware]) doStart(e Hardware) {
	pi, _ := e.(THardware)
	a.start.Start(pi)
}

func (a *appWrapper[THardware]) ApplicationMainLoop() AppMainLoopFunc {
	if a.mainLoop == nil {
		return nil
	}
	return a.doMainLoop
}

func (a *appWrapper[THardware]) doMainLoop(e Hardware, deltaTime time.Duration) {
	pi, _ := e.(THardware)
	a.mainLoop.MainLoop(pi, deltaTime)
}

func (a *appWrapper[THardware]) ApplicationEnd() AppEndFunc {
	if a.end == nil {
		return nil
	}
	return a.doEnd
}

func (a *appWrapper[THardware]) doEnd(e Hardware) {
	if a.end != nil {
		pi, _ := e.(THardware)
		a.end.End(pi)
	}
}
