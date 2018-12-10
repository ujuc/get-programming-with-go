package main

import (
	"fmt"
	"strings"
)

func CharFind(char rune) rune {
	var result rune

	if char > 'Z' {
		result = (char % 26) + 'A'
	}

	return result
}

func Decipher(cipherText string, keyword string) string {
	var result []string
	keyword = strings.Repeat(keyword, len(cipherText))

	for i, c := range(cipherText) {
		if c >= 'A' && c <= 'Z' {
			c += int32(keyword[i])
			c = CharFind(c)
		}
		result = append(result, string(c))
	}
	return strings.Join(result, "")
}

func main() {
	cipherText 	:= "CSOITEUIWUIZNSROCNKFD"
	keyword 	:= "GOLANG"

	var result string = Decipher(cipherText, keyword)
	fmt.Println(result)
}
