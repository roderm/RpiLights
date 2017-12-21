package light

import (
	"github.com/stianeikeland/go-rpio"
	"time"
	"context"
	"fmt"
)

type ColorScheme struct {
	Red uint8
	Green uint8
	Blue uint8
}
type RpiLight struct {
	Frequency int64
	ctx 	context.Context
	cancel 	func()
	brightness int
	ledR 	rpio.Pin
	ledG 	rpio.Pin
	ledB 	rpio.Pin
	cs ColorScheme
}

func NewLight(pinR int, pinG int, pinB int, f int64) (RpiLight, error)  {
	err := rpio.Open()
	if err!= nil {
		return RpiLight{}, err
	}
	ledR := rpio.Pin(pinR)
	ledG := rpio.Pin(pinG)
	ledB := rpio.Pin(pinB)
	ledR.Output()
	ledG.Output()
	ledB.Output()
	return RpiLight{
		ledR:ledR,
		ledG:ledG,
		ledB:ledB,
		Frequency:f,
	cs:ColorScheme{
		Red:255,
		Green:255,
		Blue:255}}, nil
}
func (l *RpiLight) On() {
	if l.cancel != nil {
		l.cancel()
	}
	l.ctx, l.cancel = context.WithCancel(context.Background())
	go l.run()
}

func (l *RpiLight) Off() {
	if l != nil{
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

}

func (l *RpiLight) DimTo(brightness int) {
	for {
		if brightness == l.brightness {
			fmt.Println("TargetBrigtness reached")
			return
		}

		if brightness > l.brightness {
			l.brightness++
			time.Sleep(time.Millisecond *100)
		}else if brightness < l.brightness {
			l.brightness--
			time.Sleep(time.Millisecond *100)
		}
	}
	fmt.Println("Exit func?")
}
func (l *RpiLight) run(){
	go func() {
		cycle := 0
		sleeps := time.Second / time.Duration(l.Frequency*2)
		fmt.Println("Sleeps:", sleeps)
		for {
			select {
			case <-l.ctx.Done():
				rpio.Close()
				return
			default:
			// Brightness Ups
				if l.brightness <= 0 {
					l.brightness = 100
				}
			/*
				ups := (100 / (l.brightness))
				if cycle% ups == 0 {
					l.ledR.Write(rpio.High)
					l.ledG.Write(rpio.High)
					l.ledB.Write(rpio.High)
				}else{
					l.ledR.Write(rpio.Low)
					l.ledG.Write(rpio.Low)
					l.ledB.Write(rpio.Low)
				}
			*/
				// red
				if cycle*100 % (100 / (l.brightness )) == 0 {
					l.ledR.Write(rpio.High)
				}else{
					l.ledR.Write(rpio.Low)
				}
				// green
				if cycle% (100 / (l.brightness)) == 0 {
					l.ledG.Write(rpio.High)
				}else{
					l.ledG.Write(rpio.Low)
				}
				// blue
				if cycle% (100 / (l.brightness)) == 0 {
					l.ledB.Write(rpio.High)
				}else{
					l.ledB.Write(rpio.Low)
				}


				cycle++
				if cycle >= int(l.Frequency*2){
					cycle = 0
				}
				time.Sleep(sleeps)
			}
		}
	}()
}
