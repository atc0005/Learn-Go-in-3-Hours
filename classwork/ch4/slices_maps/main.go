package main

import "fmt"

func main() {

	// https://blog.golang.org/go-maps-in-action
	//
	// var m map[string]int
	//
	// Map types are reference types, like pointers or slices, and so the
	// value of m above is nil; it doesn't point to an initialized map. A nil
	// map behaves like an empty map when reading, but attempts to write to a
	// nil map will cause a runtime panic; don't do that. To initialize a map,
	// use the built in make function:
	m := make(map[string]int)
	m["hello"] = 300

	// store the value that the 'hello' key references
	h := m["hello"]
	fmt.Println("hello in m via 'h' variable:", h)

	// this hasn't been set, so the zero value for this non-existent key in
	// the map is used
	fmt.Println("a in m:", m["a"])

	// referred to as the "comma ok" idiom
	// * 'v' contains the value (if present)
	// * 'ok' contains the result of attempting to retrieve the value of the
	//   specified key and is an indication of whether the key exists in the map
	if v, ok := m["hello"]; ok {
		fmt.Println("hello in m:", v)
	}

	// overwrite the value already stored in the map via the 'hello' key
	m["hello"] = 20
	fmt.Println("hello in m:", m["hello"])
}
