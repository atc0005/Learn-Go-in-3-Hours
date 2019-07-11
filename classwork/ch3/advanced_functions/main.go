package main

import "fmt"

func main() {

	// anonymous function (declaring function within a function is otherwise
	// not allowed)
	myAddOne := func(a int) int {
		return a + 1
	}

	fmt.Println(myAddOne(1))

}
