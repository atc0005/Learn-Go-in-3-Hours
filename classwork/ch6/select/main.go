package main

import (
	"fmt"
	"time"
)

func multiples(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				fmt.Printf("Current value is %d before post-increment.\n", curVal)
				curVal++

			// wait for a value to come down this channel; this channel is
			// skipped as it is not usable until a value comes down the `done`
			// channel
			case <-done:
				fmt.Println("goroutine shutting down")
				return
			}
		}
	}()

	// return a channel named `out` that the goroutine will use to pass values
	// down for use late and a separtate channel named `done` that is used to
	// control when the goroutine will end
	return out, done
}

func main() {
	twosCh, done := multiples(2)

	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}

	// signal our goroutine that it should stop
	close(done)
	//
	// Q: How quickly does it stop? Are we waiting for the "random" case for
	// reading the `done` channel to hit or is it immediate?
	//
	// A: Brief testing shows that the curVal variable is incremented once
	// more and sent down the channel (11), but is never read, so the rest
	// of this application never sees it.

	//do more stuff
	time.Sleep(1 * time.Second)
}
