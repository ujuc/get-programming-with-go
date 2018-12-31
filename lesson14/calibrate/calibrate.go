package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		fmt.Printf("%v' K\n", s())
		time.Sleep(time.Second)
	}
}

func fakeSensor() sensor {
	return func() kelvin {
		return kelvin(rand.Intn(151) + 150)
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(realSensor, offset)
	fmt.Println(sensor())

	measureTemperature(5, fakeSensor())
}
