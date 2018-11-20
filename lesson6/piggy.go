package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var contain float64 = 0.0

	for contain <= 20.00 {
		rand.Seed(time.Now().UnixNano())

		switch num := rand.Intn(3); num {
		case 0:
			contain += 0.10
		case 1:
			contain += 0.25
		default:
			contain += 0.05
		}

		fmt.Printf("%4.2f\n", contain)
	}
}
