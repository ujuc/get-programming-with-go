package main

import "fmt"


func rot13(message string) string {
	var decode string

	for _, c := range message {
		if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		} else if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		decode += string(c)
	}
	return decode
}

func main() {
	message := "Hola Estación Espacial Internacional"
	var encode string

	for _, c := range message {
		if c >= 'A' && c <= 'Z' {
			c = c - 13
			if c < 'A' {
				c = c + 26
			}
		} else if c >= 'a' && c <= 'z' {
			c = c - 13
			if c < 'a' {
				c = c + 26
			}
		}
		encode += string(c)
	}

	fmt.Println("Encode: ", encode)

	decode := rot13(encode)
	fmt.Println("Decode: ", decode)
}
