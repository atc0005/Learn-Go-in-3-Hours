package main

import "fmt"

func main() {

	// when declared this way, the variable is assigned the zero value for a
	// pointer: nil
	var b *int

	fmt.Println(b)
	// <nil>

	fmt.Println(*b)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x491190]
	//
	// goroutine 1 [running]:
	// main.main()
	//         Learn-Go-in-3-Hours/classwork/ch3/pointers/main.go:12 +0x80
	// exit status 2
}
