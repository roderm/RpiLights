package main

//go:generate protoc --go_out=plugins=grpc:. grpc/RpiLight.proto

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"playground/light"
	"playground/server"
	"syscall"
)

func main() {
	light.Setup(2, 3, 4, 40000)
	mlight := light.GetLight()
	fmt.Println("Light created")
	mserver := server.TelnetSever{
		Light: mlight}
	fmt.Println("Start serving")
	cancel, err := mserver.Serve(6600, context.Background())
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	func() {
		<-c
		cancel()
		os.Exit(1)
	}()
}
