package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(celsius(k.celsius()).fahrenheit())
}

func main() {
	var k kelvin = 294.0
	var c celsius = 127.0
	fmt.Println("K -> F", k.fahrenheit())
	fmt.Println("K -> C", k.celsius())
	fmt.Println("C -> F", c.fahrenheit())
}
