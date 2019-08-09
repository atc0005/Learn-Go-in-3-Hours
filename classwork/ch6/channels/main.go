package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	// Write to the `in` channel twice before attempting to read from it
	// in the goroutine. This causes a panic:
	// fatal error: all goroutines are asleep - deadlock!
	in <- 1
	//	in <- 2
	o1 := <-out

	// The workaround is to make at least one of the channels buffered OR
	// sequence the actions so that we write to `in`, then read from `out`
	// before writing once more to `in` and so on.
	in <- 2
	o2 := <-out
	fmt.Println(o1, o2)
}
