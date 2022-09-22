// Clockwerk
// author: Adam Wonak (github.com/awonak)
// date: 2022-09-12
// labels: clock
//
// Clock multiplier / divider with a range of /16 to *8 from the main clock speed.
//
// Note: Clocks will run stable within the 20-240 bpm range, however when
// several clocks are at audio rate and set to different multiplications, the
// clocks may become unstable.
//
// digital_in: unused (TODO: enable external clock)
// analog_in: unused
//
// knob_1: adjust clock tempo between 20 and 240 BPM
// knob_2: adjust the multiplication/division factor of the selected clock output
//
// button_1: move the selected clock param edit to the left
// button_2: move the selected clock param edit to the right
//
// dual_press: reset all clocks to resync
//
// output_[1-6]: clock output [1-6]
package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/tinydraw"

	europi "github.com/awonak/EuroPiGo"
)

const (
	MinBPM     = 20
	MaxBPM     = 240
	DefaultBPM = 120
	PPQN       = 4
	ResetDelay = time.Duration(2 * time.Second)
)

var (
	// Positive values are multiplications and negative values are divisions.
	DefaultFactor = [6]int{1, 2, 4, -2, -4, -8}
	FactorChoices []int
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
	*europi.EuroPi

	bpm      int
	clocks   [6]int
	resets   [6]chan int
	selected int
	external bool
	period   time.Duration

	displayShouldUpdate bool
	clocksShouldReset   bool
	lastClockChange     time.Time
}

func (c *Clockwerk) editParams() {
	_bpm := c.readBPM()
	if _bpm != c.bpm {
		c.bpm = _bpm
		c.displayShouldUpdate = true
		c.clocksShouldReset = true && !c.external
		c.lastClockChange = time.Now()
	}

	_factor := c.readFactor()
	if _factor != c.clocks[c.selected] {
		c.clocks[c.selected] = _factor
		c.displayShouldUpdate = true
		c.clocksShouldReset = true
		c.lastClockChange = time.Now()
	}
}

func (c *Clockwerk) readBPM() int {
	// Provide a range of 19 - 240 bpm. bpm < 20 will switch to external clock.
	_bpm := c.K1.Range((MaxBPM+1)-(MinBPM-2)) + MinBPM - 1
	if _bpm < MinBPM && c.period > 0 {
		_bpm = int((time.Minute)/(c.period*PPQN)) + 1
		c.external = true
	} else {
		c.external = false
	}
	return _bpm
}

func (c *Clockwerk) readFactor() int {
	return FactorChoices[c.K2.Range(len(FactorChoices))]
}

func (c *Clockwerk) startClocks() {
	for i := 0; i < len(c.clocks); i++ {
		c.resets[i] = make(chan int)
		go c.clock(i, c.resets[i])
	}
}

func (c *Clockwerk) stopClocks() {
	for _, c := range c.resets {
		c <- 0
	}
}

func (c *Clockwerk) resetClocks() {
	c.stopClocks()
	c.startClocks()
}

func (c *Clockwerk) clock(i int, reset chan int) {
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

		high, low := c.clockPulseWidth(c.clocks[i])

		c.CV[i].On()
		t = t.Add(high)
		time.Sleep(t.Sub(time.Now()))

		c.CV[i].Off()
		t = t.Add(low)
		time.Sleep(t.Sub(time.Now()))
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
	c.Display.WriteLine(fmt.Sprintf("%1sBPM: %3d", external, c.bpm), 2, 8)
	// c.Display.WriteLine(fmt.Sprintf("PW: 50%%"), europi.OLEDWidth/2, 8)

	// Display each clock multiplication or division setting.
	for i, factor := range c.clocks {
		text := " 1"
		switch {
		case factor < -1:
			text = fmt.Sprintf("\\%d", -factor)
		case factor > 1:
			text = fmt.Sprintf("x%d", factor)
		}
		c.Display.WriteLine(fmt.Sprintf("%-3v", text), int16(i*europi.OLEDWidth/len(c.clocks))+2, 26)
	}
	xWidth := int16(europi.OLEDWidth / len(c.clocks))
	xOffset := int16(c.selected) * xWidth
	// TODO: replace box with chevron.
	tinydraw.Rectangle(c.Display, xOffset, 16, xWidth, 16, europi.White)

	if c.clocksShouldReset {
		tinydraw.Rectangle(c.Display, 0, 0, 128, 32, europi.White)
	}

	c.Display.Display()
}

func main() {
	c := Clockwerk{
		EuroPi:              europi.New(),
		clocks:              DefaultFactor,
		displayShouldUpdate: true,
		clocksShouldReset:   false,
	}

	// Lower range value can have lower sample size
	c.K1.Samples(500)
	c.K2.Samples(20)

	c.DI.Handler(func(pin machine.Pin) {
		// Measure current period between clock pulses.
		c.period = time.Now().Sub(c.DI.LastInput())
	})

	// Move clock config option to the left.
	c.B1.Handler(func(p machine.Pin) {
		if c.B2.Value() {
			c.clocksShouldReset = true
			return
		}
		c.selected = europi.Clamp(c.selected-1, 0, len(c.clocks))
		c.displayShouldUpdate = true
	})

	// Move clock config option to the right.
	c.B2.Handler(func(p machine.Pin) {
		if c.B1.Value() {
			c.clocksShouldReset = true
			return
		}
		c.selected = europi.Clamp(c.selected+1, 0, len(c.clocks)-1)
		c.displayShouldUpdate = true
	})

	// Init parameter configs based on current knob positions.
	c.bpm = c.readBPM()
	c.clocks[c.selected] = c.readFactor()

	c.startClocks()

	for {
		// Check for clock updates every 2 seconds.
		time.Sleep(ResetDelay)
		lastChange := time.Now().Sub(c.lastClockChange)
		if c.clocksShouldReset && lastChange > ResetDelay {
			c.resetClocks()
			c.clocksShouldReset = false
			c.displayShouldUpdate = true
		}
	}
}
