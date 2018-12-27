package main

import "fmt"

// kelvin 에서 celsius 로 변환한다.
func kelvinToCelsius (k float64) float64 {
	k -= 273.15
	return k
}

// celsius 에서 fahrenheit 로 변환한다.
func celsiusToFahrenheit (c float64) float64 {
	f := ((c * 9.0) / 5.0) + 32.0
	return f
}

// kelvin 에서 fahrenheit 로 변환한다.
func kelvinToFahrenheit (k float64) float64 {
	c := kelvinToCelsius(k)
	f := celsiusToFahrenheit(c)
	return f
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "'K is ", celsius, "'C")

	kelvin = 0.0
	fahrenheit := kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "'K is ", fahrenheit, "'F")
}
