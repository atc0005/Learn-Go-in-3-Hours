package main

import (
	"fmt"
	"os"
)

func main() {

	word := os.Args[1]
	c := "crackerjack"

	// here we are initializing the 'l' variable, but are not specifying a
	// variable for comparison (leaving that to each case statement)
	switch l := len(word); {

	// Note: The first case statement that returns true is executed

	case word == "hi":
		fmt.Println("Very informal!")
		fallthrough
	case word == "hello":
		fmt.Println("Hi yourself")
	case l == 1:
		fmt.Println("I don't know any one letter words")
	case 1 < l && l < 10, word == c:
		fmt.Println("This word is either", c, "or it is 2-9 characters long")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long")
	}

}
