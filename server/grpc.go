package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "rpilight/grpc"
	"rpilight/light"
)

var instance Service

type Service struct {
	light *light.RpiLight
}

func StartServer(port int) (*Service, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	serv := GetService()
	pb.RegisterRpiLightServer(s, serv)
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return serv, nil

}
func GetService() *Service {
	if &instance == nil {
		instance = Service{}
	}
	return &instance
}

func (s *Service) SetLight(light *light.RpiLight) {
	s.light = light
}

func (s *Service) On(ctx context.Context, n *pb.Empty) (*pb.Empty, error) {
	s.light.On()
	return &pb.Empty{}, nil
}

func (s *Service) Off(ctx context.Context, n *pb.Empty) (*pb.Empty, error) {
	s.light.Off()
	return &pb.Empty{}, nil
}

func (s *Service) GetInfo(ctx context.Context, n *pb.Empty) (*pb.State, error) {
	colors := s.light.GetColors()
	state := pb.State{
		State:  s.light.GetState(),
		Colors: &colors,
		Bright: &pb.Brightness{
			Value: int32(s.light.GetBrightness())}}
	return &state, nil
}

func (s *Service) SetColor(ctx context.Context, n *pb.ColorScheme) (*pb.Empty, error) {
	cs := pb.ColorScheme{
		Red:   n.Red,
		Green: n.Green,
		Blue:  n.Blue}
	s.light.SetColors(cs)
	return &pb.Empty{}, nil
}
func (s *Service) SetBrightness(ctx context.Context, n *pb.Brightness) (*pb.Empty, error) {
	b := int(n.Value)
	s.light.SetBrightness(b)
	return &pb.Empty{}, nil
}
func (s *Service) SubscribeStateChange(n *pb.Empty, stream pb.RpiLight_SubscribeStateChangeServer) error {
	colors := s.light.GetColors()
	state := pb.State{
		State:  s.light.GetState(),
		Colors: &colors,
		Bright: &pb.Brightness{
			Value: int32(s.light.GetBrightness())}}
	stream.Send(&state)
	c := make(chan pb.State)
	s.light.RegisterChannel(c)
	defer close(c)
	for {
		state := <-c
		stream.Send(&state)
	}
	return nil
}
