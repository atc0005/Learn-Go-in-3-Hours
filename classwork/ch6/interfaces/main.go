package main

import "fmt"

func main() {

	// Empty interface; ANY type in Go matches it.
	// The empty interface is like the "void" pointer in C/C++ or the "object"
	// type in Java.
	// This is Go's way of saying, "This could be anything".
	//
	// Go provides two ways to get back to the concrete type behind an
	// interface; there's no way to do anything else with an empty interface.
	// 1) Type assertion
	// 2) Type switch
	var i interface{}

	i = "Hello"

	// Type assertion. You can only do a type assertion on an interface. You
	// can do a type conversion on any type at all.
	j := i.(string)

	// If you specify an additional variable (which is "ok" by
	// convention), ok will be assigned the value of true if the assertion
	// worked and false if it didn't. If the assertion worked the first
	// variable specified will get the value from the interface. If the
	// assertion fails it will be assigned the zero value of that type.
	k, ok := i.(int)
	fmt.Println(j, k, ok)
	// Hello 0 false

	// panic: interface conversion: interface {} is string, not int
	m := i.(int)
	fmt.Println(m)
}
