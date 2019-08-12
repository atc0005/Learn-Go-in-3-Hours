package main

import (
	"fmt"
	"sync"
	"time"
)

// Limit the total number of "connections" in order to provide
// stable/consistent service handling

func main() {
	workers := 10000

	// create a buffered channel of functions that take an int and returns an
	// int. The buffer size is set to the number of specified workers.
	pool := make(chan func(int) int, workers)
	for i := 0; i < workers; i++ {

		// for each iteration, a function is put into the pool
		pool <- func(in int) int {
			// NOTE: Add intentional delay here to emulate a service unable to
			// keep up with demand
			time.Sleep(5 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	count := 0
	totalStart := time.Now()
	for i := 0; i < 100000; i++ {
		start := time.Now()

		// try to get a worker from the pool
		select {

		// if successful, launch a goroutine, run the worker, then return the
		// worker to the pool.
		//
		// Said another way, we try to read a "value" from the `pool` channel
		// (which in this case is a function that receives an int and returns
		// an int), save the function to `f`, then use `f` to perform a task
		// before we send `f` back into the channel for reading in the next
		// interation. The default case statement is used whenever a worker is
		// not available (e.g., channel is empty, blocked from reading), or in
		// other words, whenever we run so many goroutines that we cannot
		// finish running them and returning them to keep the pool from
		// running empty.
		case f := <-pool:
			fmt.Printf("ACCEPT | Workers available: %d\n", len(pool))
			fmt.Println("processing", i)
			count++
			wg.Add(1)
			go func(in int) {
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				pool <- f
				wg.Done()
			}(i)

		// reject the request if there is not an available worker
		default:
			fmt.Printf("REJECT | Workers available: %d\n", len(pool))
			fmt.Println("rejecting request", i, "too busy")
		}
	}
	wg.Wait()
	fmt.Println("total processed:", count)
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
