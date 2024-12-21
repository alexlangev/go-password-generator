package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	pwGen(3, true)
	pwGen(19, false)
	pwGen(15, true)
}

func pwGen (length int, incSpecChars bool) string {
	regChars := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	specChars := []rune{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*','+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}

	pw := ""

	charSpace := make([]rune, len(regChars))
	copy(charSpace, regChars)
	if incSpecChars {
		charSpace = append(charSpace, specChars...)
	}

	for i := 0; i < length; i++ {
		pw += string(charSpace[rand.IntN(len(charSpace))])
	}

	fmt.Println(pw)
	return pw
}
