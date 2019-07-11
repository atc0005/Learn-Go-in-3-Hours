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

// receive a function (which takes an int and returns an int), return a
// function (that takes an int and returns an int)
func makeDoubler(f func(int) int) func(int) int {

	// return our closure
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {

	addOne := makeAdder(1)
	doubleAddOne := makeDoubler(addOne)

	// fn := makeAdder(1)
	// fmt.Println(fn(2))

	fmt.Println(addOne(1))
	fmt.Println(doubleAddOne(1))

}
