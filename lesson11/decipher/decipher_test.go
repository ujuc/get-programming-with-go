package main

import (
	"testing"
)

func TestDecipherV1(t *testing.T) {
	cipherText := "C"
	keyword := "G"

	v1 := Decipher(cipherText, keyword)
	if v1 != "I" {
		t.Fatal("Decipher(C, G) == I")
	}

	cipherText 	= "CSOITE"
	keyword 	= "GOLANG"

	v2 := Decipher(cipherText, keyword)
	if v2 != "IGZIGK" {
		t.Fatal("Decipher(CSOITE, GOLANG) == IGZIGK")
	}

	cipherText	= "CSOITEUIWU"
	keyword		= "GOLANG"

	v3 := Decipher(cipherText, keyword)
	if v3 != "IGZIGKAWHU" {
		t.Fatal("Decipher(CSOITEUIWU, GOLANG) == IGZIGKAWHU")
	}
}

func TestCharFind(t *testing.T) {
	v1 := CharFind('C' + 'G')
	if v1 != ('I') {
		t.Fatal("CharFind('C' + 'G') == 'I'")
	}

	v2 := CharFind('S' + 'O')
	if v2 != ('G') {
		t.Fatal("CharFind('S' + 'O') == 'G'")
	}
}
