package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "playground/grpc"
)

func main() {
	conn, err := grpc.Dial("192.168.178.95:6600", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewRpiLightClient(conn)

	stream, err := client.SubscribeStateChange(context.Background(), &pb.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		in, err := stream.Recv()
		log.Println(in)
		if err == io.EOF {
			panic(io.EOF)
		}
		if err != nil {
			panic(err)
		}
	}
}
