package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	// since this channel is NOT buffered we would be deadlocked here after
	// writing to the channel (while we wait for the value to be read from the
	// channel)
	// fatal error: all goroutines are asleep - deadlock!
	//out <- 1
	//
	// Never reached!
	//fmt.Println("Made it past sending 1 on the int typed `in` channel")

	// A Select statement looks a lot like a switch statement, only the cases
	// are reads and writes to channels. If any case can succeed that read or
	// write happens and the commands for that case are executed. If we run
	// this program now we see that we can skip over the case for `in` and go
	// straight to the case for `out`, then we print out the message that we
	// got from `out`.
	//
	// If more than one case in the Select statement can be run, because
	// multiple channels are ready to be read or written, the Select statement
	// picks one case at random to run. It does not go in order. By ensuring
	// random behavior this prevents certain kinds of deadlocks that can
	// happen in concurrent systems.
	select {

	// write 2 to the `in` channel.
	// Blocking writes cause this case to be skipped and the next one
	// evaluated. To make this case eligible for use, we have to make sure the
	// channel is buffered. In this example, it is not.
	case in <- 2:
		fmt.Println("wrote 2 to in")
		//fmt.Println(<-in)
	case v := <-out:
		fmt.Println("read", v, "from out")

	// In this example, this is the ONLY eligible case statement due to the
	// other unbuffered channels blocking
	default:
		fmt.Println("nothing else works")
	}

}
