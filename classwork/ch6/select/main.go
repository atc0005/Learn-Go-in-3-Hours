package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate allowing more connections to service than we can handle. Later
// example uses back-pressure to gracefully handle that connection overload.

func main() {
	var wg sync.WaitGroup
	totalStart := time.Now()
	// got 49682 for 24841 after 1.153999s
	// got 50884 for 25442 after 1.151998s
	// got 50084 for 25042 after 1.153999s
	// panic: too many concurrent operations on a single file or socket (max 1048575)

	// goroutine 1073951 [running]:
	// internal/poll.(*fdMutex).rwlock(0xc000080280, 0x0, 0xc235ccfde8)
	//         c:/go/src/internal/poll/fd_mutex.go:147 +0x146
	// internal/poll.(*FD).writeLock(...)
	//         c:/go/src/internal/poll/fd_mutex.go:239
	// internal/poll.(*FD).Write(0xc000080280, 0xc25e4313c0, 0x29, 0x40, 0x0, 0x0, 0x0)
	//         c:/go/src/internal/poll/fd_windows.go:659 +0x55
	// os.(*File).write(...)
	//         c:/go/src/os/file_windows.go:224
	// os.(*File).Write(0xc00007e008, 0xc25e4313c0, 0x29, 0x40, 0x57dbc0, 0xc235ccfee0, 0x408ad2)
	//         c:/go/src/os/file.go:145 +0x77
	// fmt.Fprintln(0x4eb200, 0xc00007e008, 0xc235ccff48, 0x6, 0x6, 0x57cfe0, 0x3baa040c, 0x0)
	//         c:/go/src/fmt/print.go:266 +0x92
	// fmt.Println(...)
	//         c:/go/src/fmt/print.go:275
	// main.main.func1(0xbf4c729a85020f6c, 0xdba3eaa7d, 0x57cfe0, 0xc000052070, 0x1062fd)
	//         T:/github/Learn-Go-in-3-Hours/classwork/ch6/select/main.go:22 +0x1a7
	// created by main.main
	//         T:/github/Learn-Go-in-3-Hours/classwork/ch6/select/main.go:18 +0x101
	// exit status 2
	//for i := 0; i < 100000000; i++ {
	for i := 0; i < 100000; i++ {
		start := time.Now()
		wg.Add(1)
		go func(in int) {

			// delay for one second
			time.Sleep(1 * time.Second)

			// double the number
			out := 2 * in
			_ = out

			// calculate the time needed
			fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))

			// signal that the goroutine is finished
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
