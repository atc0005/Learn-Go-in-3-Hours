package main

import "fmt"

func main() {
	in := make(chan int, 10)
	out := make(chan int)

	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	go func() {
		for {

			// In order to tell the difference between reading a closed
			// channel or reading a zero value for a channel, put a comma and
			// a second variable, usually called `ok`, on the left side of the
			// equal sign. If `ok` is true, then the channel is still open, if
			// the value of `ok` is false the channel is closed.
			i, ok := <-in

			// as soon as `ok` evaluates to false close the channel and break
			// out of the loop
			if !ok {

				// Closed means that channel will not have any more values
				// written to it; closing a channel does not wipe out its
				// contents. If a buffered channel is closed, any values in
				// the buffer are still available to be read. If you read
				// from a closed channel with no more data the read returns
				// immediately with the zero value for the type of the
				// channel.
				close(out)
				break
			}
			out <- i * 2
		}
	}()

	// range over the channel until the channel is closed
	for v := range out {
		fmt.Println(v)
	}
}
