package main

import (
	"time"
	"fmt"
	"math/rand"
)

type kelvin float64

func function() {
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	measureTemperature(3, fakeSensor)
}

func realSensor() kelvin {
	return 0
}

func measureTemperature(samples int, sensor func() kelvin ) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v k\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}