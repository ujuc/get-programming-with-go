package main

import (
	"fmt"
	"sort"
)

func terraform(planets []string) []string {
	for i := range planets {
		switch planets[i] {
		case "Mars", "Uranus", "Neptune":
			planets[i] = "New " + planets[i]
		}
	}

	return planets;
}

func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planetsNew := terraform(planets)
	sort.StringSlice(planetsNew).Sort()
	fmt.Println(planetsNew)
}
