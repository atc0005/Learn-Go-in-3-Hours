package main

import "fmt"

func main() {

	s3 := []string{"a", "b", "c"}

	// range returns an index, and the matching value for index
	for k, v := range s3 {
		fmt.Println(k, v)
	}

	// create a slice from a slice using slice expression
	//
	// provide starting point and ending point for the slice, optionally
	// leaving off starting point if starting point is 0 or the ending point
	// if the ending point is the length of the slice
	//
	// WARNING: Both the original slice and the new slice refer to the same
	// area of memory. Changing one changes the other. Slices are reference
	// types and behave like pointers.
	s4 := s3[0:2]
	fmt.Println("s4:", s4)

	s3[0] = "d"
	fmt.Println("s4:", s4)

	// zero value for a slice is nil
	// zero value for a pointer is nil
	var s5 []string
	s5 = s3

	// Note: s5[1] is read out loud as "s 5 sub 1"
	s5[1] = "camel"
	fmt.Println("s3:", s3)

	modSlice(s3)
	fmt.Println("s3[0]:", s3[0])

}

func modSlice(s []string) {
	s[0] = "hello"
}
