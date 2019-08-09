package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

// you can make a function implement an interface by defining a function type
// ...
type testerFunc func(int) bool

// ... and a method on the function type.
func (tf testerFunc) test(i int) bool {

	// when this method is invoked we will call the underlying function, pass
	// it the value that was passed into the method, and return the value from
	// the function as the value returned from the method. Because our test
	// func type defines a method called test that takes in an int and returns
	// a bool, it implements the tester interface.
	return tf(i)
}

func main() {

	// We're calling runTests, passing the value 10 and a slice of
	// tester interface implementations
	result := runTests(10, []tester{

		// Declare an anonymous function that takes in an int and returns a
		// bool. The anonymous function is "wrapped" in testerFunc() to
		// perform a type conversion to make the anonymous function a
		// testerFunc.
		testerFunc(func(i int) bool {
			return i%2 == 0
		}),

		// Second anonymous function that's also converted to type
		// testerFunc.
		//
		// Rather than having to define a type ahead of time, like we did in
		// our first example, we were able to provide the behavior for our
		// interface on the fly by using a type conversion to make a function
		// into a type that implemented an interface.
		testerFunc(func(i int) bool {
			return i < 20
		}),
	})
	fmt.Println(result)

	// You should be aware that converting functions to interface
	// implementations is not just a corner case brainteaser, it's actually
	// incredibly useful. Remember the http.HandlerFunc function we used all
	// the way back in the first video of section 2. It took in a function and
	// used it to respond to an HTTP request. But what's interesting is that
	// Go's HTTP server registers types and implements an interface called
	// http.Handler.
	//
	// In order to make a function work as a http.Handler, Go defines a type
	// called HandlerFunc, which has a method on HandlerFunc that meets the
	// http.Handler interface. This allows http.HandlerFunc function to use a
	// type conversion to convert functions with the Write() function
	// signature into instances of type http.HandlerFunc. This allows you to
	// use a closure in your program as an instance of an interface, all
	// completely invisible to your program. So that's the power of implicit
	// interfaces, and the ability to declare methods on any user-defined
	// type.

}
