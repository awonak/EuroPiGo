package main

import (
	"fmt"
	"machine"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"tinygo.org/x/tinydraw"
)

const (
	DefaultBPM  = 120
	MinBPM      = 20
	MaxBPM      = 240
	PPQN        = 4
	MaxDivision = 16
)

var (
	DefaultDivisions = [2]int{1, 1}
	DivisionChoices  []int
)

func init() {
	// Create range of -16 to -2 for division and 1 to 16 for multiplication.
	for i := -16; i <= 16; i++ {
		if i != -1 && i != 0 {
			DivisionChoices = append(DivisionChoices, i)
		}
	}
}

type Clockwerk struct {
	europi.EuroPi

	bpm         int
	selectedDiv int
	divisions   [2]int

	displayShouldUpdate bool
	prevk2              int
}

func (c *Clockwerk) editParams() {
	_bpm := c.K1.Range(MaxBPM-MinBPM+1) + MinBPM
	if _bpm != c.bpm {
		c.bpm = _bpm
		c.displayShouldUpdate = true
	}

	i := c.K2.Range(len(DivisionChoices))
	if c.prevk2 != DivisionChoices[i] {
		c.divisions[c.selectedDiv] = DivisionChoices[i]
		c.prevk2 = DivisionChoices[i]
		c.displayShouldUpdate = true
	}
}

func (c *Clockwerk) clock(idx int) {
	t := time.Now()
	for {
		// Add expensive call to clock goroutine to factor its time into clock sleep.
		c.updateDisplay()
		c.editParams()

		// TODO: Make pulse width configurable.
		high, low := c.clockPulseWidth(c.divisions[idx])

		c.CV[idx].On()
		t = t.Add(high)
		time.Sleep(t.Sub(time.Now()))

		c.CV[idx].Off()
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
		c.Display.WriteLine(fmt.Sprintf("BPM: %3d     PW: 50%%", c.bpm), 0, 8)

		// Divisions
		for i, div := range c.divisions {
			var divText string
			switch {
			case div < -1:
				divText = fmt.Sprintf("\\%d", -div)
			case div > 1:
				divText = fmt.Sprintf("x%d", div)
			default:
				divText = " 1"
			}
			c.Display.WriteLine(fmt.Sprintf("%-3v", divText), int16(i*europi.OLEDWidth/len(c.divisions))+2, 26)
		}
		xWidth := int16(europi.OLEDWidth / len(c.divisions))
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
		divisions:           DefaultDivisions,
	}

	// Lower range value can have lower sample size
	c.K1.Samples(500)
	c.K2.Samples(20)

	c.B1.Handler(func(p machine.Pin) {
		c.selectedDiv = europi.Clamp(c.selectedDiv-1, 0, len(c.divisions))
		c.displayShouldUpdate = true
	})

	c.B2.Handler(func(p machine.Pin) {
		c.selectedDiv = europi.Clamp(c.selectedDiv+1, 0, len(c.divisions))
		c.displayShouldUpdate = true
	})

	for i := 0; i < len(c.divisions); i++ {
		go c.clock(i)
	}

	for {
		// Long sleep between goroutine executions
		time.Sleep(1 * time.Second)
	}
}
