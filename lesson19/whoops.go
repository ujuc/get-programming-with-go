package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector zz9",
		"Mars":  "Sector zz9",
	}

	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
