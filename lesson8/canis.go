package main

import "fmt"

// solution: https://en.wikipedia.org/wiki/Canis_Major_Overdensity
func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792.458
	const secondsPerDay = 86400

	const days = distance / lightSpeed / secondsPerDay
	const years = days / 365.25

	fmt.Println("Canis Major Dwarf is", years, "light years away.")
}
