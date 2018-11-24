package main

import "fmt"

func main() {
	fmt.Printf("%v is a %[1]T\n", "itieral string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}
