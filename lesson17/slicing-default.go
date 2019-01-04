package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	allPlanets := planets[:]
	fmt.Println(terrestrial, gasGiants, iceGiants, allPlanets)

	neptune := "Neptune"
	tune := neptune[3:]

	fmt.Println(tune)
}
