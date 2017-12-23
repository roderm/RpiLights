package light

import (
	"context"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

var instance RpiLight

type ColorScheme struct {
	Red   uint8
	Green uint8
	Blue  uint8
}
type RpiLight struct {
	Frequency  int64
	ctx        context.Context
	cancel     func()
	brightness int
	ledR       rpio.Pin
	ledG       rpio.Pin
	ledB       rpio.Pin
	cs         ColorScheme
}

func Setup(pinR int, pinG int, pinB int, f int64) {
	err := rpio.Open()
	if err != nil {
		panic(err)
		return
	}
	ledR := rpio.Pin(pinR)
	ledG := rpio.Pin(pinG)
	ledB := rpio.Pin(pinB)
	ledR.Output()
	ledG.Output()
	ledB.Output()
	instance = RpiLight{
		ledR:      ledR,
		ledG:      ledG,
		ledB:      ledB,
		Frequency: f,
		cs: ColorScheme{
			Red:   255,
			Green: 255,
			Blue:  255}}
}
func GetLight() *RpiLight {
	return &instance
}
func (l *RpiLight) On() {
	if l.cancel == nil {
		l.ctx, l.cancel = context.WithCancel(context.Background())
		go l.run()
	}
}

func (l *RpiLight) Off() {
	if l.cancel != nil {
		l.cancel()
		l.cancel = nil
	}
	l.ledR.Write(rpio.Low)
	l.ledG.Write(rpio.Low)
	l.ledB.Write(rpio.Low)
}

func (l *RpiLight) SetBrightness(brightness int) {
	l.brightness = brightness
}

func (l *RpiLight) SetColors(cs ColorScheme) {
	l.cs = cs
}

func (l *RpiLight) DimTo(brightness int) {
	for {
		if brightness == l.brightness {
			fmt.Println("TargetBrigtness reached")
			return
		}

		if brightness > l.brightness {
			l.brightness++
			time.Sleep(time.Millisecond * 100)
		} else if brightness < l.brightness {
			l.brightness--
			time.Sleep(time.Millisecond * 100)
		}
	}
}
func (l *RpiLight) run() {
	go func() {
		cycle := 0
		sleeps := time.Nanosecond / time.Duration(l.Frequency)
		fmt.Println("Sleeps:", sleeps)
		fmt.Println("Colors", l.cs)

		setPin := func(cycle int, ups float32, pin rpio.Pin, color uint8) {
			if color > 0 {
				br := float32(float32(255) / float32(color))
				if cycle%int(ups*br) == 0 {
					pin.Write(rpio.High)
				} else {
					pin.Write(rpio.Low)
				}
			} else {
				pin.Write(rpio.Low)
			}
		}
		for {
			select {
			case <-l.ctx.Done():
				return
			default:
				if l.brightness <= 0 {
					l.brightness = 100
				}

				ups := float32(100 / (l.brightness))
				// red
				setPin(cycle, ups, l.ledR, l.cs.Red)
				setPin(cycle, ups, l.ledG, l.cs.Green)
				setPin(cycle, ups, l.ledB, l.cs.Blue)

				cycle++
				if cycle >= int(l.Frequency) {
					cycle = 0
				}
				time.Sleep(sleeps)
			}
		}
	}()
}
