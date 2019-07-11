package main

import "fmt"

// Per Learn Go in 3 Hours course:
// Go doesn't have optional parameters, named parameters or function
// overloading

func addNumbers(a int, b int) int {
	return a + b
}

func main() {

	a := addNumbers(2, 3)
	fmt.Println(a)

	a = addNumbers(4, 10)
	fmt.Println(a)

	a = addNumbers(100, -100)
	fmt.Println(a)

}
