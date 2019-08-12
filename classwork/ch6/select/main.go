package main

import "fmt"

func multiples(i int) chan int {
	out := make(chan int)
	curVal := 0
	go func() {
		for {
			// infinite duration goroutine that continually writes an
			// incrementing value down the channel
			// Emphasis: IT NEVER ENDS
			out <- curVal * i
			curVal++
		}
		// BUG: unreachable code (per go-vet)
		fmt.Println("Is this ever reached?")
	}()

	// return the created channel that the goroutine will use to pass values
	// down for use later
	return out
}

func main() {
	twosCh := multiples(2)

	// ranging over a channel that is closed will trigger the loop to end,
	// but since the channel is not closed we need some other way to determine
	// when to exit the for range loop. Here we use the returned value: when
	// the value exceeds 20 we stop looping.
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}

	//do more stuff
}
