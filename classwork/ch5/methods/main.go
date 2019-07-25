package main

import "fmt"

type Foo struct {
	A int
	B string
}

// '(f Foo)' is a method receiver
// ...
// method receiver declared between the keyword func and the name of the
// method. The method receiver looks like a parameter declaration. It consists
// of a left parenthesis, and identifier, the type in which we are defining
// the method, and a right parenthesis. Whenever you want to access a field
// and a struct inside of this method you use this identifier. It is like the
// self the first parameter and a Python method declaration or the implicit in
// this reference that you get any Java method.
//
// Note: This is a value receiver (i.e., refer to Foo directly)
func (f Foo) String() string {
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}

// Double  does not have a return value since this method does in-place
// modification of the method receiver
//
// Note: This is a "reference" receiver (due to pointer to Foo).
func (f *Foo) Double() {
	f.A = f.A * 2
}

func main() {

	f := Foo{
		A: 10,
		B: "Hello",
	}

	// Before modification
	fmt.Println(f.String())

	// Modify struct field
	f.Double()

	// After modification
	fmt.Println(f.String())

}
