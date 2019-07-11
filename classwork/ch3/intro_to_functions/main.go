package main

import "fmt"

// Per Learn Go in 3 Hours course:
// Go doesn't have optional parameters, named parameters or function
// overloading

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func addNumbers(a int, b int) int {
	return a + b
}

func main() {

	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)

	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	divAndRemainder(-1, 20)

}
