package main

//go:generate protoc --go_out=plugins=grpc:. grpc/RpiLight.proto

import (
	"context"
	"fmt"
	"time"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/signal"
	"rpilight/light"
	"rpilight/hwlight"
	"rpilight/server"
	"syscall"
	"os/user"
	pb "rpilight/grpc"
)

type Config struct {
	TelnetPort int
	Frequency  int
	GPIORed    int
	GPIOGreen  int
	GPIOBlue   int
	GrpcPort   int
	PWMMode	   string
}

func ReadConfig(configfile string ) Config {
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
	hwlight.MyPwmTest()
	fmt.Println("Test finished")
	var conffile string
	if  len(os.Args) < 2 {
		conffile = "./config.toml"
	}else {
		conffile = os.Args[1]
	}

	conf := ReadConfig(conffile)

	mlight := createLight(conf)
	fmt.Println("Light created")
	mserver := server.TelnetSever{
		Light: mlight}

	mlight.On();
	mlight.SetColors(pb.ColorScheme{
		Red:0,
		Green:255,
		Blue:0});

	fmt.Println("Start serving (Telnet)")
	cancel, err := mserver.Serve(conf.TelnetPort, context.Background())
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second);

	mlight.SetColors(pb.ColorScheme{
		Red:255,
		Green:0,
		Blue:0});

	fmt.Println("Start serving (GRPC)")
	s := server.GetService()
	s.SetLight(mlight)
	server.StartServer(conf.GrpcPort)

	time.Sleep(time.Second);

	mlight.SetColors(pb.ColorScheme{
		Red:0,
		Green:0,
		Blue:255});
	/*fmt.Println("Starting bonjour")
	server.AdvertiseBonjour([]string{"Simple GrpcLight"}, conf.GrpcPort)
	*/

	// Show that programm has startet:

	time.Sleep(time.Second);
	mlight.SetColors(pb.ColorScheme{
		Red:255,
		Green:255,
		Blue:255});
	time.Sleep(time.Second)
	mlight.Off();

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	func() {
		<-c
		cancel()
		os.Exit(1)
	}()
}

func createLight(conf Config) light.ILight {
	if conf.PWMMode == "hardware" {
		fmt.Println("Hardware mode selected")
		hwlight.Setup(conf.GPIORed, conf.GPIOGreen, conf.GPIOBlue, int64(conf.Frequency))
		return hwlight.GetLight()
	}else {
		fmt.Println("Software mode selected")
		light.Setup(conf.GPIORed, conf.GPIOGreen, conf.GPIOBlue, int64(conf.Frequency))
		return light.GetLight()
	}
}