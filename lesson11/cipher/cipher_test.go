package main

import (
	"testing"
)

func TestCharFind(t *testing.T) {
	v1 := CharFind('Y' - 'G')
	if v1 != 'S' {
		t.Fatal("CharFind('Y' - 'G') == 'S'")
	}

	v2 := CharFind('O' - 'O')
	if v2 != 'A' {
		t.Fatal("CharFind('O' - 'O') == 'A'")
	}

	v3 := CharFind('M' - 'N')
	if v3 != 'Z' {
		t.Fatal("charFind('M' - 'N') == 'Z'")
	}
}

func TestCipher(t *testing.T) {
	plainText := "y"
	keyword := "G"

	v1 := Cipher(plainText, keyword)
	if v1 != "S" {
		t.Fatal("Cipher('y', 'G') == 'S'")
	}

	plainText = "your"
	keyword = "GOLA"

	v2 := Cipher(plainText, keyword)
	if v2 != "SAJR" {
		t.Fatal("Cipher('your', 'GOLA') == 'SAJR'")
	}

	plainText = "your message"
	keyword = "GOLANG"

	v3 := Cipher(plainText, keyword)
	if v3 != "SAJRZYMEPGR" {
		t.Fatal("Cipher('your message', 'GOLANG') == 'SAJRZYMEPGR'")
	}

	plainText = " "
	keyword = "GOLANG"

	v4 := Cipher(plainText, keyword)
	if v4 != "" {
		t.Fatal("Cipher (' ', 'GOLANG') == ''")
	}
}
