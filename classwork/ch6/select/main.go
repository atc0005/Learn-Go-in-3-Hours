package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int, 1)

	// since this channel is buffered we're not blocked trying to write to
	// the channel
	out <- 1

	// we ARE however blocked from writing to `in` since there are no other
	// goroutines around to read from the channel (reminder: `main()` is also
	// a goroutine)
	in <- 2

	// FIX: Use a `select` statement to manage channels

	fmt.Println("wrote 2 to in")
	v := <-out
	fmt.Println("read", v, "from out")
}
