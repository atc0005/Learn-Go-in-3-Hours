package main

import (
	"fmt"
)

func main() {

	// create a buffered channel that can be used (asynchronously) to hold up
	// to 10 values before writes to the channel resume the normal
	// synchronous/blocking behavior
	//
	// NOTE: Most of the time you will wish to use synchronous/unbuffered
	// channels so that you can better understand/follow data flow (e.g.,
	// "baton in a relay race")
	out := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out <- localI * 2
		}(i)
	}
	var result []int
	for i := 0; i < 10; i++ {

		// illustrating that we can read from the channel via the built-in
		// `append()` function without the intermediate variable (though this
		// is probably not near as clear as using the intermediate variable)
		//val := <-out
		result = append(result, <-out)
	}
	fmt.Println(result)
}
