package main

import (
	"fmt"
	"playground/light"
	"playground/server"
)

func main() {
	light.Setup(2, 3, 4, 40000)
	mlight := light.GetLight()
	fmt.Println("Light created")
	mserver := server.TelnetSever{
		Light: mlight}
	fmt.Println("Start serving")
	err := mserver.Serve(6600)
	if err != nil {
		panic(err)
	}
}
