package main

import (
	"fmt"
)

func main() {

	// Note: These channels are synchronous and bidirectional by default. The
	// direction implied by the name spells out the intent, not the
	// limitation. If we wish to make unidirectional channels we will need to
	// specify this when declaring the channels.
	in := make(chan string)
	out := make(chan string)

	// https://guzalexander.com/2013/12/06/golang-channels-tutorial.html
	//
	// i := make(chan int)       		// by default the capacity is 0
	// s := make(chan string, 3) 		// non-zero capacity
	// r := make(<-chan bool)          	// can only read from
	// w := make(chan<- []os.FileInfo) 	// can only write to

	go func() {

		// receive from the `in` channel
		name := <-in

		// write to the `out` channel
		out <- fmt.Sprintf("Hello, " + name)
	}()

	// write `Bob` string to the `in` channel; not only does our goroutine
	// have access to this value, but our goroutine will block/wait until a
	// message is received and saved to `name`
	in <- "Bob"

	// close the channel, signal that we are done using it.
	// NOTE: There are several "gotchas" to be aware of re closed channels
	close(in)

	// receive string data back from the `out` channel; we closed the `in`
	// channel only at this point.
	//
	// NOTE: Leaving a channel open is not considerd bad practice, but
	// presumably (not really sure yet) this applies mostly to channels that
	// we are done sending data down from the main code body?
	//
	// NOTE: A WaitGroup is not needed to block early exit as Go will wait
	// for a message from the `out` channel until one is received.
	message := <-out

	fmt.Println(message)
}
