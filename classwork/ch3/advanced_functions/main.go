package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

// receive function as second argument. This second argument can be any
// function that takes one integer argument and returns one integer
func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func main() {

	printOperation(1, addOne)
	printOperation(1, addTwo)

}
