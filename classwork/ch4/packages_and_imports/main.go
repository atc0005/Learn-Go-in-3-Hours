package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "👋 🌍"

	// Print the string
	fmt.Println(s)

	// Print the length in bytes
	fmt.Println(len(s))

	// Print the number of "runes"
	fmt.Println(utf8.RuneCountInString(s))

}
