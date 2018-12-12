package main

import (
	"fmt"
	"strings"
)

func CharFind(char rune) rune {
	var result rune

	char = char % 26

	if char >= 0 {
		result = char + 'A'
	} else if char < 0 {
		result = char + 1 + 'Z'
	}
	return result
}

func Cipher(plainText string, keyword string) string {
	var result []string

	plainText = strings.ToUpper(plainText)
	plainText = strings.Replace(plainText, " ", "", -1)
	keyword = strings.Repeat(keyword, len(plainText))

	for i, c := range(plainText) {
		if c >= 'A' && c <= 'Z' {
			c -= int32(keyword[i])
			c = CharFind(c)
		}
		result = append(result, string(c))
	}

	return strings.Join(result, "")
}

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"

	var result string = Cipher(plainText, keyword)
	fmt.Println(result)
}
