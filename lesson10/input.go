package main

import (
	"fmt"
	"strconv"
)

func boolChecker(i int) (string, error) {
	if i < 0 {
		return "", fmt.Errorf("음수(%v)는 nono.", i)
	}

	if i > 1 {
		return "", fmt.Errorf("only 0, 1 ok?")
	}

	return strconv.Itoa(i), nil
}

func main() {
	input_data, err := boolChecker(0)
	if err != nil {
		fmt.Println(err)
	}

	switch input_data {
	case "true", "yes", "1":
		fmt.Println("true")
	case "false", "no", "0":
		fmt.Println("false")
	}
}
