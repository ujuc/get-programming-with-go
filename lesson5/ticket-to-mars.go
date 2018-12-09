package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func rangeRandom(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	fmt.Printf("%-18v %-4v %-10v %-5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println(strings.Repeat("=", 40))

	rand.Seed(time.Now().UnixNano())
	for count := 0; count < 10; count++ {
		spaceline := ""
		days := 0
		trip_type := ""
		price := 0

		switch rand.Intn(3) {
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		default:
			spaceline = "Space Adventures"
		}

		mars_away := 62100000
		speed := rangeRandom(16, 30)
		days = mars_away / (speed * 60 * 60 * 24)

		switch speed {
		case 16, 17, 18, 19, 20:
			price = rangeRandom(36, 40)
		case 21, 22, 23, 24, 25:
			price = rangeRandom(40, 45)
		case 26, 27, 28, 29, 30:
			price = rangeRandom(46, 50)
		}

		if rand.Intn(2) == 0 {
			trip_type = "One-way"
		} else {
			trip_type = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-18v %4v %-10v $%4v\n", spaceline, days, trip_type, price)
	}
}
