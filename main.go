package main

import (
	"playground/light"
	"log"
	"playground/server"
	"fmt"
)

func main() {
	mlight, err := light.NewLight(2,3,4,40000)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Light created")
	mserver := server.TelnetSever{
		Light:mlight}
	fmt.Println("Start serving")
	err = mserver.Serve(6600)
	if err != nil {
		panic(err)
	}
}