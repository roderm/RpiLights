package light

import (
	pb "rpilight/grpc"
)

type ILight interface {
	On()
	Off()
	GetState() pb.LightState
	GetBrightness() int
	GetColors() pb.ColorScheme
	SetColors(pb.ColorScheme)
	SetBrightness(int)
	DimTo(int)
	RegisterChannel(chan pb.State)
}
