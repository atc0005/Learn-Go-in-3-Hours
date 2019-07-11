package main

import "fmt"

func main() {
	uniHello := "👋 🌍"

	// my understanding: convert/cast copy of string to slice of bytes
	// video: copy that string into a slice of bytes
	bytes := []byte(uniHello)

	// print slice of bytes (as numeric values, not a "string")
	fmt.Println(bytes)

	// copy our string into a slice of runes
	runes := []rune(uniHello)

	// print slice of runes
	fmt.Println(runes)

	// NOTE: A slice of runes provides random access to the runes in a string

	// this modifies a copy, not the original string
	runes[1] = 'a'
	fmt.Println(runes)
	fmt.Println(uniHello)
}
