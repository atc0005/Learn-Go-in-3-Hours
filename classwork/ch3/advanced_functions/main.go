package main

import "fmt"

func main() {

	b := 2

	myAddOne := func(a int) {

		// write to 'b' variable from main body of main()
		b = a + b
	}

	myAddOne(1)
	fmt.Println(b)

}
