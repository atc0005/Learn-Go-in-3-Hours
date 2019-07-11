package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func main() {

	// not invoking, referencing the original function
	myAddOne := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1))

}
