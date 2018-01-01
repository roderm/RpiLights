package light

import (
	"context"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"

	pb "rpilight/grpc"
)

var instance RpiLight

type RpiLight struct {
	Frequency      int64
	ctx            context.Context
	cancel         func()
	brightness     int
	ledR           rpio.Pin
	ledG           rpio.Pin
	ledB           rpio.Pin
	cs             pb.ColorScheme
	NotifyChannels []chan pb.State
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
		cs: pb.ColorScheme{
			Red:   255,
			Green: 255,
			Blue:  255}}
}
func GetLight() ILight {
	return instance
}

func (l RpiLight) GetState() pb.LightState {
	if l.cancel == nil {
		return pb.LightState_OFF
	}
	if l.cancel != nil {
		return pb.LightState_ON
	}
	return pb.LightState_UNKNOWN
}
func (l RpiLight) On() {
	if l.cancel == nil {
		l.ctx, l.cancel = context.WithCancel(context.Background())
		go l.run()
	}
	l.triggerStateChange()
}

func (l RpiLight) Off() {
	if l.cancel != nil {
		l.cancel()
		l.cancel = nil
		// To be sure loop has finished
		time.Sleep(time.Millisecond * 50)
	}
	l.ledR.Write(rpio.Low)
	l.ledG.Write(rpio.Low)
	l.ledB.Write(rpio.Low)
	l.triggerStateChange()
}

func (l RpiLight) SetBrightness(brightness int) {
	l.brightness = brightness
	l.triggerStateChange()
}
func (l RpiLight) GetBrightness() int {
	return l.brightness
}
func (l RpiLight) SetColors(cs pb.ColorScheme) {
	l.cs = cs
	l.triggerStateChange()
}
func (l RpiLight) GetColors() pb.ColorScheme {
	return l.cs
}

func (l RpiLight) DimTo(brightness int) {
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
func (l RpiLight) run() {
	go func() {
		cycle := 0
		sleeps := time.Nanosecond / time.Duration(l.Frequency)
		fmt.Println("Sleeps:", sleeps)
		fmt.Println("Colors", l.cs)

		setPin := func(cycle int, ups float32, pin rpio.Pin, color int32) {
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
func (l RpiLight) RegisterChannel(c chan pb.State) {
	l.NotifyChannels = append(l.NotifyChannels, c)
}
func (l RpiLight) triggerStateChange() {
	state := pb.State{
		State:  l.GetState(),
		Colors: &l.cs,
		Bright: &pb.Brightness{
			Value: int32(l.GetBrightness())}}

	go func() {
		for _, c := range l.NotifyChannels {
			c <- state
		}
	}()
}
