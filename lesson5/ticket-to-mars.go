package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func range_random(min, max int) int {
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

		switch num := rand.Intn(3); num {
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		default:
			spaceline = "Space Adventures"
		}

		mars_away := 62100000
		speed := range_random(16, 30)
		days = mars_away / (speed * 60 * 60 * 24)

		switch speed {
		case 16, 17, 18, 19, 20:
			price = range_random(36, 40)
		case 21, 22, 23, 24, 25:
			price = range_random(40, 45)
		case 26, 27, 28, 29, 30:
			price = range_random(46, 50)
		}

		if num := rand.Intn(2); num == 0 {
			trip_type = "One-way"
		} else {
			trip_type = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-18v %4v %-10v $%4v\n", spaceline, days, trip_type, price)
	}
}
