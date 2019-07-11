package main

import "fmt"

func main() {

	// arg1: type of slice to make
	// arg2: slice length
	// arg3: optional capacity for preallocating a slice (optimization)
	s := make([]string, 0)
	fmt.Println("length of s:", len(s))

	// extend slice by using built-in append() function
	s = append(s, "hello")
	fmt.Println("length of s:", len(s))
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0])

	// starts off with length of 2
	s2 := make([]string, 2)
	fmt.Println("contents of s2[0]:", s2[0])

	// You have to be careful when calling append() on a slice that already
	// has a length. append() is not going to put a value into the first
	// element in s2. It's going to increase the length of the slice and add a
	// new element to the end.
	//
	// Q: Is this because the first two elements already have "zero" values?
	s2 = append(s2, "hello")
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))

	// Extend example above to help make it clearer that first two elements
	// have empty strings.
	fmt.Printf("contents of s2[0]: %q\n", s2[0])
	fmt.Printf("contents of s2[1]: %q\n", s2[1])
	fmt.Printf("contents of s2[2]: %q\n", s2[2])

}
