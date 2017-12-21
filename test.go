package main

import (
	"log"
	"time"
	"playground/light"
)

func test()  {
	mlight, err := light.NewLight(2,3,4,40000)
	if err != nil {
		log.Panic(err)
	}

	mlight.SetBrightness(100)
	mlight.On()

	/*
	r := []int{10, 20, 30, 40, 50, 60, 100, 90, 80, 60, 70, 40, 20, 5, 4, 3, 2, 1}
	for i := range r {
		mlight.SetBrightness(r[i])
		fmt.Println(fmt.Sprint("d%", r[i]))
		time.Sleep(500 * time.Millisecond)
	}
	*/
	mlight.DimTo(100)
	mlight.SetColors(light.ColorScheme{
		Red:1,
		Green:1,
		Blue:1})
	time.Sleep(2 * time.Second)
	mlight.Off()
}
