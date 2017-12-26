package main

//go:generate protoc --go_out=plugins=grpc:. grpc/RpiLight.proto

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/signal"
	"playground/light"
	"playground/server"
	"syscall"
)

type Config struct {
	TelnetPort int
	Frequency  int
	GPIORed    int
	GPIOGreen  int
	GPIOBlue   int
	GrpcPort   int
}

func ReadConfig() Config {
	var configfile = "./config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
func main() {
	conf := ReadConfig()
	light.Setup(conf.GPIORed, conf.GPIOGreen, conf.GPIOBlue, int64(conf.Frequency))
	mlight := light.GetLight()
	fmt.Println("Light created")
	mserver := server.TelnetSever{
		Light: mlight}
	fmt.Println("Start serving (Telnet)")
	cancel, err := mserver.Serve(conf.TelnetPort, context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Start serving (GRPC)")
	s := server.GetService()
	s.SetLight(mlight)
	server.StartServer(conf.GrpcPort)

	/*fmt.Println("Starting bonjour")
	server.AdvertiseBonjour([]string{"Simple GrpcLight"}, conf.GrpcPort)
	*/
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	func() {
		<-c
		cancel()
		os.Exit(1)
	}()
}
