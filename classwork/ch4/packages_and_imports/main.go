package main

import (
	"fmt"
	"strings"
)

// type rune = int32
//     rune is an alias for int32 and is equivalent to int32 in all ways. It is
//     used, by convention, to distinguish character values from integer values.
func rot13(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return ((((in - 'A') + 13) % 26) + 'A')
	}
	if in >= 'a' && in <= 'z' {
		return ((((in - 'a') + 13) % 26) + 'a')
	}

	// Return characters outside of A-Z and a-z as-is without modification
	return in
}

func main() {

	// Original string
	s := "This is a test 123 😃"

	// Modified version of string
	s2 := strings.Map(rot13, s)
	fmt.Println(s2)

	// Modified version of string converted back to original form
	s3 := strings.Map(rot13, s2)
	fmt.Println(s3)
}
