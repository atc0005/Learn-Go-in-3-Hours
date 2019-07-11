package main

import "fmt"

func main() {

	b := 2

	myAddOne := func(a int) int {

		// read 'b' variable from main body of main()
		return a + b
	}

	fmt.Println(myAddOne(1))

}
