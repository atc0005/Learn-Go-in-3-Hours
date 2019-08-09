package main

import (
	"fmt"
)

func main() {

	// create a buffered channel that can be used (asynchronously) to hold up
	// to 10 values before writes to the channel resume the normal
	// synchronous/blocking behavior
	out := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out <- localI * 2
		}(i)
	}
	var result []int
	for i := 0; i < 10; i++ {
		val := <-out
		result = append(result, val)
	}
	fmt.Println(result)
}
