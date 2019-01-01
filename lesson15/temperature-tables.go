package main

import (
	"fmt"
	"strings"
)

type celsius float64
type fahrenheit float64
type temperature string

func drawTable(t temperature) {
	switch t {
	case "c":
		fmt.Println(strings.Repeat("=", 23))
		fmt.Printf("| %-8v | %-8v |\n", "'C", "'F")
		fmt.Println(strings.Repeat("=", 23))

		var step celsius
		for step = -40.0; step <= 100.0; step += 5 {
			var c = step
			fmt.Printf("| %-8v | %-8v |\n", c, c.fahrenheit())
		}
		fmt.Println(strings.Repeat("=", 23))
		break
	case "f":
		fmt.Println(strings.Repeat("=", 23))
		fmt.Printf("| %-8v | %-8v |\n", "'F", "'C")
		fmt.Println(strings.Repeat("=", 23))

		var step fahrenheit
		for step = -40.0; step <= 100.0; step += 5 {
			var f = step
			fmt.Printf("| %-8v | %-8.2f |\n", f, f.celsius())
		}
		fmt.Println(strings.Repeat("=", 23))
	}
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}


func main() {
	drawTable("c")
	drawTable("f")
}
