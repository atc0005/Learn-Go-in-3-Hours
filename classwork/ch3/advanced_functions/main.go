package main

import "fmt"

func main() {

	b := 2

	// Closure: local function that has access to the variables that exist in
	// the environment in which it was declared
	myAddOne := func(a int) {

		// write to 'b' variable from main body of main()
		b = a + b
	}

	myAddOne(1)
	fmt.Println(b)

}
