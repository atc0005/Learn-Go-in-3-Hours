package main

import "fmt"

// receive an int, return an anonymous function (which takes an int and
// returns an int). This is a closure; the original value that was passed in
// is kept as part of the state for later use.
//
// e.g., `fn := makeAddr(1); fmt.Println(fn(2))` results in three being printed
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func main() {

	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fn := makeAdder(1)
	fmt.Println(fn(2))

	fmt.Println(addOne(1))
	fmt.Println(addTwo(1))

}
