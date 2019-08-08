package main

import (
	"fmt"
)

func reallyNil() error {
	var e error
	fmt.Println("e is nil:", e == nil)

	// Q: Is this safe because 'error' is a base type?
	return e
}

type MyError struct {
	A       int
	B       int
	Message string
}

// Declare a method Error() on a pointer to MyError; this method matches
// the interface signature for Error, which makes a pointer to MyError
// implement the Error interface.
func (me *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error %s", me.A, me.B, me.Message)
}

// While using your own error type is a good practice, you need to be careful.
// The Go error handling pattern always compares the return error value
// against nil. In order to understand how to properly return a custom error
// from a function requires you to understand how nil works with interfaces.
// The zero value for all interfaces is nil, but this nil is a bit different
// from the nil for pointers, slices, or maps. An interface is considered nil
// only if it isn't associated with any underlying value, even an underlying
// nil value.

func notReallyNil() error {
	var me *MyError
	fmt.Println("me is nil:", me == nil)

	// See notes below; this is not a good idea as it associates an underlying
	// type with the error interface
	return me
}

func main() {

	e := reallyNil()
	me := notReallyNil()
	// The first two print statements tell you that e is nil and me is nil.
	// This is as expected because we have var e error and we have var me
	// *MyError, and in both cases you have a var without assigning a value to
	// it, that gives you the zero value. The zero value for the pointer here
	// is nil.
	//
	// However, when we assign reallyNil() back to e, `e == nil`
	// is true, but `me == nil` is false.  This is because an interface is
	// only nil if there's no underlying type assigned to it.
	//
	// Even though the value pointed to by the concrete type MyError is nil,
	// what we returned back from notReallyNil() was a nil MyError, and there
	// was still an underlying type associated with an error interface that
	// was returned.
	//
	// So therefore the interface me in main is not actually nil.
	//
	// Now be aware this behavior isn't specific to the error interface. It
	// applies to all interface types, and this is probably the most
	// counterintuitive thing in Go, and personally I consider it a mistake in
	// the language.
	//
	// If you are defining your own error types you need to be sure that you
	// never define a variable to be your own error type, because if you
	// return that variable it will not appear nil even if no error actually
	// occurred in your function.

	fmt.Println("in main, e is nil:", e == nil)
	fmt.Println("in main, me is nil:", me == nil)
}
