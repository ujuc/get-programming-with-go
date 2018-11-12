package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var my_num = 20
	var rand_num = 100

	for {
		var guess = rand.Intn(rand_num)

		if my_num < guess {
			fmt.Println(my_num, guess, "Down")
			rand_num = rand_num - guess/2
			fmt.Println(rand_num)
		} else if my_num > guess {
			fmt.Println(my_num, guess, "UP")
			rand_num = rand_num + guess/2
			fmt.Println(rand_num)
		} else {
			fmt.Println(my_num, guess, "mach")
			break
		}
	}
}
