package main

import "fmt"

func main() {
	a := 10
	b := &a // pointer or reference to a
	c := a
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	// Example output (pointer address subject to change):
	//
	// 10 0xc000056058 10 10
	// 30 0xc000056058 30 10
	// 30 0xc000056058 30 40

}
