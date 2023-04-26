// Clockwerk
// author: Adam Wonak (github.com/awonak)
// date: 2022-09-12
// labels: clock
//
// Clock multiplier and divider with a range of /16 to x8 from the main clock speed.
//
// Note: Clocks will run stable within the 60-240 bpm range, however when
// several clocks are at audio rate and set to different multiplications, the
// clocks may become unstable.
//
// Note: When the clock is slow and there are several divisions, the async sleep
// functions will cause the restart to complete much slower because it has to
// wait for all clocks to wake up before it can receive the restart signal.
//
// - digital_in: external clock
// - analog_in: unused
//
// - knob_1: adjust clock tempo between 60 and 240 BPM
// - knob_2: adjust the multiplication/division factor of the selected clock output
//
// - button_1: move the selected clock param edit to the left
// - button_2: move the selected clock param edit to the right
//
// - dual_press: reset all clocks to resync
//
// - output_[1-6]: clock output [1-6]
package main

import (
	"strconv"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont/proggy"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/lerp"
)

const (
	MinBPM     = 60
	MaxBPM     = 240
	DefaultBPM = 120
	PPQN       = 4
	ResetDelay = time.Duration(2 * time.Second)
)

var (
	// Positive values are multiplications and negative values are divisions.
	DefaultFactor = [6]int{1, 2, 4, -2, -4, -8}
	FactorChoices []int
	DefaultFont   = &proggy.TinySZ8pt7b
)

func init() {
	// Seed the clock multiplication and division options.
	// TODO: fix pulsewidth period for these divisions.
	// DivisionChoices = append(DivisionChoices, -64)
	// DivisionChoices = append(DivisionChoices, -32)
	// Create range of -16 to -2 for division and 1 to 8 for multiplication.
	for i := -16; i <= 8; i++ {
		if i != -1 && i != 0 {
			FactorChoices = append(FactorChoices, i)
		}
	}
}

type Clockwerk struct {
	displayShouldUpdate bool
	doClockReset        bool
	external            bool
	selected            uint8
	bpm                 uint16
	prevk2              int
	period              time.Duration
	clocks              [6]int
	resets              [6]chan uint8
	writer              fontwriter.Writer

	*europi.EuroPi
	bpmLerp    lerp.Lerper32[uint16]
	factorLerp lerp.Lerper32[int]
}

func (c *Clockwerk) editParams() {
	_bpm := c.readBPM()
	if _bpm != c.bpm {
		c.bpm = _bpm
		c.displayShouldUpdate = true
	}

	_factor := c.readFactor()
	if _factor != c.prevk2 {
		c.clocks[c.selected] = _factor
		c.prevk2 = _factor
		c.displayShouldUpdate = true
	}
}

func (c *Clockwerk) readBPM() uint16 {
	// Provide a range of 59 - 240 bpm. bpm < 60 will switch to external clock.
	_bpm := c.bpmLerp.ClampedLerpRound(c.K1.Percent())
	if _bpm < MinBPM {
		c.external = true
		_bpm = 0
		if c.period > 0 {
			_bpm = uint16((time.Minute)/(c.period*PPQN)) + 1
		}
	} else {
		c.external = false
		c.period = 0
	}
	return _bpm
}

func (c *Clockwerk) readFactor() int {
	idx := c.factorLerp.ClampedLerpRound(c.K2.Percent())
	return FactorChoices[idx]
}

func (c *Clockwerk) startClocks() {
	for i := 0; i < len(c.clocks); i++ {
		c.resets[i] = make(chan uint8)
		go c.clock(uint8(i), c.resets[i])
	}
}

func (c *Clockwerk) stopClocks() {
	for _, r := range c.resets {
		r <- 0
	}
}

func (c *Clockwerk) resetClocks() {
	c.stopClocks()
	c.startClocks()
}

func (c *Clockwerk) clock(i uint8, reset chan uint8) {
	t := time.Now()
	for {
		// Check if a reset signal has been received on the channel.
		select {
		case <-reset:
			return
		default:
		}

		// Add expensive call to clock goroutine to factor its time into clock sleep.
		c.editParams()
		c.updateDisplay()

		// External clock selected but not receiving pulses.
		if c.bpm == 0 {
			continue
		}

		high, low := c.clockPulseWidth(c.clocks[i])

		c.CV[i].SetCV(1.0)
		t = t.Add(high)
		time.Sleep(time.Since(t))

		c.CV[i].SetCV(0.0)
		t = t.Add(low)
		time.Sleep(time.Since(t))
	}
}

func (c *Clockwerk) clockPulseWidth(factor int) (high, low time.Duration) {
	var period int
	if factor < 1 {
		// Clock divisions increase sleep duration time.
		period = int(c.sleepPeriod()) * -factor
	} else {
		// Clock multiplications decreace sleep time.
		period = int(c.sleepPeriod()) / factor
	}
	// TODO: use configuratble pulse width.
	return time.Duration(period) / 2, time.Duration(period) / 2
}

// sleepPeriod returns the duration of a ppqn clock sleep period for bpm.
func (c *Clockwerk) sleepPeriod() time.Duration {
	return time.Duration(uint64(time.Minute) / uint64(c.bpm) / PPQN)
}

func (c *Clockwerk) updateDisplay() {
	if !c.displayShouldUpdate {
		return
	}
	c.displayShouldUpdate = false
	c.Display.ClearBuffer()

	// Master clock and pulse width.
	var external string
	if c.external {
		external = "^"
	}
	c.writer.WriteLine(external+"BPM: "+strconv.Itoa(int(c.bpm)), 2, 8, draw.White)

	// Display each clock multiplication or division setting.
	dispWidth, _ := c.Display.Size()
	divWidth := int(dispWidth) / len(c.clocks)
	for i, factor := range c.clocks {
		text := " 1"
		switch {
		case factor < -1:
			text = "\\" + strconv.Itoa(-factor)
		case factor > 1:
			text = "x" + strconv.Itoa(factor)
		}
		c.writer.WriteLine(text, int16(i*divWidth)+2, 26, draw.White)
	}
	xWidth := int16(divWidth)
	xOffset := int16(c.selected) * xWidth
	// TODO: replace box with chevron.
	_ = tinydraw.Rectangle(c.Display, xOffset, 16, xWidth, 16, draw.White)

	_ = c.Display.Display()
}

var app Clockwerk

func startLoop(e *europi.EuroPi) {
	app.EuroPi = e
	app.clocks = DefaultFactor
	app.displayShouldUpdate = true
	app.writer = fontwriter.Writer{
		Display: e.Display,
		Font:    DefaultFont,
	}
	app.bpmLerp = lerp.NewLerp32[uint16](MinBPM-1, MaxBPM)
	app.factorLerp = lerp.NewLerp32(0, len(FactorChoices)-1)

	// Lower range value can have lower sample size
	_ = app.K1.Configure(hal.AnalogInputConfig{
		Samples: 500,
	})
	_ = app.K2.Configure(hal.AnalogInputConfig{
		Samples: 20,
	})

	app.DI.Handler(func(_ bool, deltaTime time.Duration) {
		// Measure current period between clock pulses.
		app.period = deltaTime
	})

	// Move clock config option to the left.
	app.B1.Handler(func(_ bool, deltaTime time.Duration) {
		if app.B2.Value() {
			app.doClockReset = true
			return
		}
		app.selected = uint8(clamp.Clamp(int(app.selected)-1, 0, len(app.clocks)))
		app.displayShouldUpdate = true
	})

	// Move clock config option to the right.
	app.B2.Handler(func(_ bool, deltaTime time.Duration) {
		if app.B1.Value() {
			app.doClockReset = true
			return
		}
		app.selected = uint8(clamp.Clamp(int(app.selected)+1, 0, len(app.clocks)-1))
		app.displayShouldUpdate = true
	})

	// Init parameter configs based on current knob positions.
	app.bpm = app.readBPM()
	app.prevk2 = app.readFactor()

	app.startClocks()
}

func mainLoop() {
	if app.doClockReset {
		app.doClockReset = false
		app.resetClocks()
		app.displayShouldUpdate = true
	}
	europi.DebugMemoryUsage()
}

func main() {
	startLoop(europi.New())

	// Check for clock updates every 2 seconds.
	ticker := time.NewTicker(ResetDelay)
	defer ticker.Stop()
	for {
		<-ticker.C
		mainLoop()
	}
}
