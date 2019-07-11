package main

import "fmt"

func setTo10(a *int) {

	// allocate new memory for a pointer and set a to reference that memory
	// location in place of the memory location provided to setTo10 as a
	// function parameter
	a = new(int)

	// Instead of changing the value stored in the original memory location
	// (i.e., changing the original variable value), we are changing the value
	// stored at the location created by our new of `new(int)`
	*a = 10
}

func main() {

	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)

}
