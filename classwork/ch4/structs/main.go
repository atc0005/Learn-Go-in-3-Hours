package main

import "fmt"

type Foo struct {

	// due to uppercase leading character, this is "public"
	// Go encourages public fields for simple data structures
	A int
	b string
}

func main() {

	f := Foo{
		A: 20,
	}

	// declare, initialize to zero values
	var f2 Foo

	// copy values into 'f2' struct from 'f' struct
	// Note: The struct type is NOT a reference type, so values are copied by
	// value and not by reference. This results in a copy of all values from
	// 'f' into 'f2', leaving the structs pointed at different memory locations
	f2 = f

	// Because the structs are referring to entirely different memory locations,
	// this makes it safe to modify one struct without concern of modifying
	// the other struct unintentionally.
	f2.A = 100

	// prove that the values for each struct are distinct
	fmt.Println(f2.A)
	fmt.Println(f.A)

	// create pointer to 'f' struct ...
	//var f3 *Foo = &f
	//
	// Note: This works, but go-lint doesn't like it:
	//
	// should omit type *Foo from declaration of var f3; it will be inferred
	// from the right-hand side
	//
	// Shorthand declaration:
	f3 := &f

	//  ... by which we can directly modify the original struct's public
	// field 'A' to a new value
	f3.A = 200

	// Show that the values are NOT distinct
	fmt.Println(f3.A)
	fmt.Println(f.A)

}
