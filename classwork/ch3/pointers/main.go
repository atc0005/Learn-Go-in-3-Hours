package main

import "fmt"

func main() {

	// when declared this way, the variable is assigned the zero value for a
	// pointer: nil
	var b *int

	fmt.Println(b)
	// <nil>

	//fmt.Println(*b)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x491190]
	//
	// goroutine 1 [running]:
	// main.main()
	//         Learn-Go-in-3-Hours/classwork/ch3/pointers/main.go:12 +0x80
	// exit status 2

	// -----------------------------------------------------------------------

	// Create a pointer, but also allocate memory for the pointer as well to
	// prevent dereferencing a nil pointer.
	// NOTE: Use of new() is discouraged by Go community
	//
	// Official doc details:
	//
	// func new(Type) *Type
	// The new built-in function allocates memory. The first argument is a type,
	// not a value, and the value returned is a pointer to a newly allocated zero
	// value of that type.
	c := new(int)

	fmt.Println(c)
	// 0xc000056058
	// Note: Memory address for the pointer

	fmt.Println(*c)
	// 0
	// Note: Zero value for pointer declared using new()
}
