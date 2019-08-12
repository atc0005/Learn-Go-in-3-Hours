package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	in2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
		wg.Done()
	}()

	go func() {
		for i := 100; i < 110; i++ {
			in2 <- i
		}
		close(in2)
		wg.Done()
	}()

	go func() {
		count := 0

		// `count` is incremented inside each case statement after the
		// associated channel is closed
		for count < 2 {

			// these case statements are only chosen if the associated channel
			// is still open
			select {

			// use `comma, ok` idiom to check if the channel is closed
			case i, ok := <-in:

				// if the channel is closed ...
				if !ok {
					fmt.Println("in2 channel is:", ok)
					count++

					// the `in` channel is set to nil; while a closed
					// channel returns a value immediately, a nil channel
					// never returns. By setting this to nil, this effectively
					// disables this case in the select from further use
					in = nil
					fmt.Println("in2 channel is:", ok)
					continue
				}

				// if the channel is NOT closed ...
				fmt.Println("in channel is:", ok)
				fmt.Println("from in, result is", i*2)

			case i, ok := <-in2:
				if !ok {
					fmt.Println("in2 channel is:", ok)
					count++
					in2 = nil
					fmt.Println("in2 channel is:", ok)
					continue
				}
				fmt.Println("in2 channel is:", ok)
				fmt.Println("from in2, result is", i+2)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
