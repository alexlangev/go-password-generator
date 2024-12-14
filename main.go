package main

import (
	"fmt"
	"math/rand/v2"
)

var regChars = [52]rune{
	'A', 
	'B', 
	'C', 
	'D', 
	'E', 
	'F', 
	'G', 
	'H', 
	'I', 
	'J', 
	'K', 
	'L', 
	'M', 
	'N', 
	'O', 
	'P', 
	'Q', 
	'R', 
	'S', 
	'T', 
	'U', 
	'V', 
	'W', 
	'X', 
	'Y', 
	'Z',
	'a', 
	'b', 
	'c', 
	'd', 
	'e', 
	'f', 
	'g', 
	'h', 
	'i', 
	'j', 
	'k', 
	'l', 
	'm', 
	'n', 
	'o', 
	'p', 
	'q', 
	'r', 
	's', 
	't', 
	'u', 
	'v', 
	'w', 
	'x', 
	'y', 
	'z',
}

var specChars = [32]rune{	
	'!', 
	'"', 
	'#', 
	'$', 
	'%', 
	'&', 
	'\'', 
	'(', 
	')', 
	'*', 
	'+', 
	',', 
	'-', 
	'.', 
	'/', 
	':', 
	';', 
	'<', 
	'=', 
	'>', 
	'?', 
	'@', 
	'[', 
	'\\', 
	']', 
	'^', 
	'_', 
	'`', 
	'{', 
	'|', 
	'}', 
	'~',
}

func main() {
	pwGen(3)
	pwGen(19)
	pwGen(15)
}

func pwGen (length int) string {
	pw := ""

	for i := 0; i < length; i++ {
		pw += string(regChars[rand.IntN(len(regChars))])
	}

	fmt.Println(pw)
	return pw
}

