package hwlight

import (
	"context"
	"fmt"
	"github.com/roderm/bcm2835"
	pb "rpilight/grpc"
	"rpilight/light"
)

var instance RpiLight

type RpiLight struct {
	Frequency      int64
	ctx            context.Context
	cancel         func()
	brightness     int
	ledR           int
	ledG           int
	ledB           int
	cs             pb.ColorScheme
	NotifyChannels []chan pb.State
}

func Setup(pinR int, pinG int, pinB int, f int64) {
	if err := bcm2835.Init(); err != nil {
		fmt.Println(err)
		return
	}
	bcm2835.GpioFsel(pinR, bcm2835.GpioFselAlt5)
	bcm2835.GpioFsel(pinG, bcm2835.GpioFselAlt5)
	bcm2835.GpioFsel(pinB, bcm2835.GpioFselAlt5)
	bcm2835.PwmSetClockDivider(bcm2835.PwmClockDivider16)
	// Channel 0
	bcm2835.PwmSetMode(0, 1, 1)
	bcm2835.PwmSetRange(0, 1024)
	// Channel 1
	bcm2835.PwmSetMode(1, 1, 1)
	bcm2835.PwmSetRange(1, 1024)
	fmt.Println("PWM-Channels are setted up")
}

func GetLight() light.ILight {
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
	fmt.Println("Switch light on")
	l.DimTo(100)
}
func (l RpiLight) Off() {
	fmt.Println("Switch light off")
	l.DimTo(0)
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
	r := (100 / 1024) * brightness
	bcm2835.PwmSetData(0, uint32(r))
	bcm2835.PwmSetData(1, uint32(r))
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
