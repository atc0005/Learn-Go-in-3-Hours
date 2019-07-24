package main

import "fmt"

type Foo struct {

	// due to uppercase leading character, this is "public"
	// Go encourages public fields for simple data structures
	A int
	b string
}

func main() {

	// declare uninitialized struct Foo
	// fields are left at "zero" values
	f := Foo{}
	fmt.Println(f)

	// declare Foo struct, initilize using positional arguments
	// Note: When using this format, ALL fields have to be specified
	g := Foo{10, "Hello"}
	fmt.Println(g)

	// declare Foo struct, initialize just 'b' field, leaving 'A' field at
	// zero value. This initialization format is more common.
	h := Foo{
		b: "Goodbye",
	}
	fmt.Println(h)

	// Access public field, set to new value
	h.A = 1000
	fmt.Println(h.A)

	// Show 'h' struct to visually compare against earlier values
	fmt.Println(h)

}
