package main

import (
	"fmt"
	"machine"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"tinygo.org/x/tinydraw"
)

const (
	MinBPM     = 20
	MaxBPM     = 240
	DefaultBPM = 120
	PPQN       = 4
)

var (
	DefaultDivisions = [6]int{1, 2, 4, -2, -4, -8}
	DivisionChoices  []int
)

func init() {
	// Seed the clock multiplication and division options.
	// DivisionChoices = append(DivisionChoices, -64)
	// DivisionChoices = append(DivisionChoices, -32)
	// Create range of -16 to -2 for division and 1 to 8 for multiplication.
	for i := -16; i <= 8; i++ {
		if i != -1 && i != 0 {
			DivisionChoices = append(DivisionChoices, i)
		}
	}
}

type Clockwerk struct {
	europi.EuroPi

	bpm         int
	selectedDiv int
	clocks      [6]int
	resets      [6]chan int

	displayShouldUpdate bool
	clocksShouldRestart bool
	prevk2              int
}

func (c *Clockwerk) editParams() {
	_bpm := c.K1.Range(MaxBPM-MinBPM+1) + MinBPM
	if _bpm != c.bpm {
		c.bpm = _bpm
		c.displayShouldUpdate = true
		c.clocksShouldRestart = true
	}

	i := c.K2.Range(len(DivisionChoices))
	if c.prevk2 != DivisionChoices[i] {
		c.clocks[c.selectedDiv] = DivisionChoices[i]
		c.prevk2 = DivisionChoices[i]
		c.displayShouldUpdate = true
		c.clocksShouldRestart = true
	}
}

func (c *Clockwerk) startClocks() {
	for i := 0; i < len(c.clocks); i++ {
		c.resets[i] = make(chan int)
		high, low := c.clockPulseWidth(c.clocks[i])
		go c.clock(c.CV[i], high, low, c.resets[i])
	}
}

func (c *Clockwerk) resetClocks() {
	for _, c := range c.resets {
		c <- 0
	}
	c.startClocks()
}

func (c *Clockwerk) clock(cv europi.Outputer, high, low time.Duration, reset chan int) {
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

		cv.On()
		t = t.Add(high)
		time.Sleep(t.Sub(time.Now()))

		cv.Off()
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
	if c.displayShouldUpdate {
		c.Display.ClearBuffer()

		// Master clock and pulse width.
		c.Display.WriteLine(fmt.Sprintf("BPM: %3d", c.bpm), 0, 8)
		c.Display.WriteLine(fmt.Sprintf("PW: 50%%"), europi.OLEDWidth/2, 8)

		// Display each clock multiplication or division setting.
		for i, div := range c.clocks {
			var divText string
			switch {
			case div < -1:
				divText = fmt.Sprintf("\\%d", -div)
			case div > 1:
				divText = fmt.Sprintf("x%d", div)
			default:
				divText = " 1"
			}
			c.Display.WriteLine(fmt.Sprintf("%-3v", divText), int16(i*europi.OLEDWidth/len(c.clocks))+2, 26)
		}
		xWidth := int16(europi.OLEDWidth / len(c.clocks))
		xOffset := int16(c.selectedDiv) * xWidth
		// TODO: replace box with chevron.
		tinydraw.Rectangle(c.Display, xOffset, 16, xWidth, 16, europi.White)

		c.Display.Display()
		c.displayShouldUpdate = false
	}
}

func main() {
	c := Clockwerk{
		EuroPi:              europi.New(),
		bpm:                 DefaultBPM,
		displayShouldUpdate: true,
		clocks:              DefaultDivisions,
	}

	// Lower range value can have lower sample size
	c.K1.Samples(500)
	c.K2.Samples(20)

	// Move clock config option to the left.
	c.B1.Handler(func(p machine.Pin) {
		c.selectedDiv = europi.Clamp(c.selectedDiv-1, 0, len(c.clocks))
		c.displayShouldUpdate = true
	})

	// Move clock config option to the right.
	c.B2.Handler(func(p machine.Pin) {
		c.selectedDiv = europi.Clamp(c.selectedDiv+1, 0, len(c.clocks)-1)
		c.displayShouldUpdate = true
	})

	c.startClocks()

	for {
		// Check for clock updates every 2 seconds.
		if c.clocksShouldRestart {
			c.resetClocks()
			c.clocksShouldRestart = false
		}
		time.Sleep(2 * time.Second)
	}
}
