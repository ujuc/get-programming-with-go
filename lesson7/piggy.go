package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var contain float64 = 0.0

	for contain <= 5.0 {
		rand.Seed(time.Now().UnixNano())

		var cent int = 0
		switch num := rand.Intn(3); num {
		case 0:
			cent += 5
		case 1:
			cent += 10
		default:
			cent += 25
		}

		contain += (float64(cent) / 100.0)

		if contain > 1.0 {
			fmt.Printf("%4.2f\n", contain)
		}
	}
}
