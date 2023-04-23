//go:build pico
// +build pico

package rev1

import (
	"fmt"
	"image/color"
	"machine"
	"math"
	"math/rand"
	"runtime/interrupt"
	"runtime/volatile"

	"github.com/heucuva/europi/internal/hardware/hal"
	"tinygo.org/x/drivers/ssd1306"
)

//============= ADC =============//

type picoAdc struct {
	adc machine.ADC
}

func newPicoAdc(pin machine.Pin) adcProvider {
	adc := &picoAdc{
		adc: machine.ADC{Pin: pin},
	}
	adc.adc.Configure(machine.ADCConfig{})
	return adc
}

func (a *picoAdc) Get(samples int) uint16 {
	if samples == 0 {
		return 0
	}

	var sum int
	state := interrupt.Disable()
	for i := 0; i < samples; i++ {
		sum += int(a.adc.Get())
	}
	interrupt.Restore(state)
	return uint16(sum / samples)
}

//============= DigitalReader =============//

type picoDigitalReader struct {
	pin machine.Pin
}

func newPicoDigitalReader(pin machine.Pin) digitalReaderProvider {
	dr := &picoDigitalReader{
		pin: pin,
	}
	dr.pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return dr
}

func (d *picoDigitalReader) Get() bool {
	state := interrupt.Disable()
	// Invert signal to match expected behavior.
	v := !d.pin.Get()
	interrupt.Restore(state)
	return v
}

func (d *picoDigitalReader) SetHandler(changes hal.ChangeFlags, handler func()) {
	pinChange := d.convertChangeFlags(changes)

	state := interrupt.Disable()
	d.pin.SetInterrupt(pinChange, func(machine.Pin) {
		handler()
	})
	interrupt.Restore(state)
}

func (d *picoDigitalReader) convertChangeFlags(changes hal.ChangeFlags) machine.PinChange {
	var pinChange machine.PinChange
	if (changes & hal.ChangeRising) != 0 {
		pinChange |= machine.PinFalling
	}
	if (changes & hal.ChangeFalling) != 0 {
		pinChange |= machine.PinRising
	}
	return pinChange
}

//============= PWM =============//

type picoPwm struct {
	pwm pwmGroup
	pin machine.Pin
	ch  uint8
	v   uint32
}

// pwmGroup is an interface for interacting with a machine.pwmGroup
type pwmGroup interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Top() uint32
	SetTop(top uint32)
	Get(channel uint8) uint32
	Set(channel uint8, value uint32)
	SetPeriod(period uint64) error
}

func newPicoPwm(pwm pwmGroup, pin machine.Pin) pwmProvider {
	p := &picoPwm{
		pwm: pwm,
		pin: pin,
	}
	return p
}

func (p *picoPwm) Configure(config hal.VoltageOutputConfig) error {
	state := interrupt.Disable()
	defer interrupt.Restore(state)

	err := p.pwm.Configure(machine.PWMConfig{
		Period: uint64(config.Period.Nanoseconds()),
	})
	if err != nil {
		return fmt.Errorf("pwm Configure error: %w", err)
	}

	p.pwm.SetTop(uint32(config.Top))
	ch, err := p.pwm.Channel(p.pin)
	if err != nil {
		return fmt.Errorf("pwm Channel error: %w", err)
	}
	p.ch = ch

	return nil
}

func (p *picoPwm) Set(v float32, ofs uint16) {
	invertedV := v * float32(p.pwm.Top())
	// volts := (float32(o.pwm.Top()) - invertedCv) - o.ofs
	volts := invertedV - float32(ofs)
	state := interrupt.Disable()
	p.pwm.Set(p.ch, uint32(volts))
	interrupt.Restore(state)
	volatile.StoreUint32(&p.v, math.Float32bits(v))
}

func (p *picoPwm) Get() float32 {
	return math.Float32frombits(volatile.LoadUint32(&p.v))
}

//============= Display =============//

const (
	oledFreq   = machine.KHz * 400
	oledAddr   = ssd1306.Address_128_32
	oledWidth  = 128
	oledHeight = 32
)

type picoDisplayOutput struct {
	dev ssd1306.Device
}

func newPicoDisplayOutput(channel *machine.I2C, sdaPin, sclPin machine.Pin) displayProvider {
	channel.Configure(machine.I2CConfig{
		Frequency: oledFreq,
		SDA:       sdaPin,
		SCL:       sclPin,
	})

	display := ssd1306.NewI2C(channel)
	display.Configure(ssd1306.Config{
		Address: oledAddr,
		Width:   oledWidth,
		Height:  oledHeight,
	})

	dp := &picoDisplayOutput{
		dev: display,
	}

	return dp
}

func (d *picoDisplayOutput) ClearBuffer() {
	d.dev.ClearBuffer()
}

func (d *picoDisplayOutput) Size() (x, y int16) {
	return d.dev.Size()
}
func (d *picoDisplayOutput) SetPixel(x, y int16, c color.RGBA) {
	d.dev.SetPixel(x, y, c)
}

func (d *picoDisplayOutput) Display() error {
	return d.dev.Display()
}

//============= RND =============//

type picoRnd struct{}

func (r *picoRnd) Configure(config hal.RandomGeneratorConfig) error {
	xl, _ := machine.GetRNG()
	xh, _ := machine.GetRNG()
	x := int64(xh)<<32 | int64(xl)
	rand.Seed(x)
	return nil
}

//============= Init =============//

func init() {
	machine.InitADC()

	RevisionMarker = newRevisionMarker()
	InputDigital1 = newDigitalInput(newPicoDigitalReader(machine.GPIO22))
	InputAnalog1 = newAnalogInput(newPicoAdc(machine.ADC0))
	OutputDisplay1 = newDisplayOutput(newPicoDisplayOutput(machine.I2C0, machine.GPIO0, machine.GPIO1))
	InputButton1 = newDigitalInput(newPicoDigitalReader(machine.GPIO4))
	InputButton2 = newDigitalInput(newPicoDigitalReader(machine.GPIO5))
	InputKnob1 = newAnalogInput(newPicoAdc(machine.ADC1))
	InputKnob2 = newAnalogInput(newPicoAdc(machine.ADC2))
	OutputVoltage1 = newVoltageOuput(newPicoPwm(machine.PWM2, machine.GPIO21))
	OutputVoltage2 = newVoltageOuput(newPicoPwm(machine.PWM2, machine.GPIO20))
	OutputVoltage3 = newVoltageOuput(newPicoPwm(machine.PWM0, machine.GPIO16))
	OutputVoltage4 = newVoltageOuput(newPicoPwm(machine.PWM0, machine.GPIO17))
	OutputVoltage5 = newVoltageOuput(newPicoPwm(machine.PWM1, machine.GPIO18))
	OutputVoltage6 = newVoltageOuput(newPicoPwm(machine.PWM1, machine.GPIO19))
	DeviceRandomGenerator1 = newRandomGenerator(&picoRnd{})
}
