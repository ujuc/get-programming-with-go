package main

import (
	"fmt"
	"strings"
)

type temperature string

var drawTable = func(t1 float64, t2 float64) {
	fmt.Printf("| %8.2f | %8.2f |\n", t1, t2)
}

func main() {
	title := func (t temperature) {
		switch t {
		case "c":
			fmt.Println(strings.Repeat("=", 23))
			fmt.Printf("| %8v | %8v |\n", "'C", "'F")
			fmt.Println(strings.Repeat("=", 23))
			break
		case "f":
			fmt.Println(strings.Repeat("=", 23))
			fmt.Printf("| %8v | %8v |\n", "'F", "'C")
			fmt.Println(strings.Repeat("=", 23))
			break
		}
	}

	bottom := func() {
		fmt.Println(strings.Repeat("=", 23))
	}

	f := func(c float64) float64 {
		return (c * 9.0 / 5.0) + 32.0
	}

	c := func(f float64) float64 {
		return (f - 32.0) * 5.0 / 9.0
	}

	title("c")
	for step := -40.0; step <= 100.0; step += 5 {
		drawTable(step, f(step))
	}
	bottom()

	title("f")
	for step := -40.0; step <= 100.0; step += 5 {
		drawTable(step, c(step))
	}
	bottom()
}
