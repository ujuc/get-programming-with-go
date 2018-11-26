package main

import (
	"fmt"
	"strconv"
)


func second() {
	countdown := 9

	str := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)
}

func third() {
	countdown, err := strconv.Atoi("10")
	if err != nil {
		//oh no, something went wrong
	}
	fmt.Println(countdown)
}

//func fourth() {
//	// constant 0.5 truncated to integer
//	var countdown = 10
//	countdown = 0.5
//	countdown = fmt.Sprintf("%v seconds", countdown)
//}

func main() {
	countdown := 10

	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	second()
	third()
	//fourth()
}
