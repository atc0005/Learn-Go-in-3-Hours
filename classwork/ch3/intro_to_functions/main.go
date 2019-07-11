package main

// Per Learn Go in 3 Hours course:
// Go doesn't have optional parameters, named parameters or function
// overloading

import "fmt"

func addNumbers(a int, b int) {
	fmt.Println(a + b)
}

func main() {

	addNumbers(2, 3)
	addNumbers(4, 10)
	addNumbers(100, -100)

}
