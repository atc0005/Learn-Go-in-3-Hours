package main

import "fmt"

func modMap(m map[string]int) {
	m["cheese"] = 20
}

func main() {

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	// declared, but not initialized?
	// i.e., safe to read, but not safe to write?
	// https://blog.golang.org/go-maps-in-action
	var m3 map[string]int

	fmt.Println("goodbye in m:", m["goodbye"])

	// maps are reference types, so assigning one map to another makes them
	// both point to the same values in memory
	m3 = m

	// this sets the value for the "goodbye" key for both m and m3
	m3["goodbye"] = 400

	// prove that we're seeing the same content
	fmt.Println("goodbye in m3:", m3["goodbye"])
	fmt.Println("goodbye in m:", m["goodbye"])

	modMap(m)
	fmt.Println("cheese in m:", m["cheese"])

}
