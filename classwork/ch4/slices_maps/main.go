package main

import "fmt"

func main() {

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50,
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// show that the key/value is present
	fmt.Println("b in m2:", m2["b"])

	// remove the key
	delete(m2, "b")

	// show that the key is no longer present
	// note: zero value is returned
	// presumably "comma ok" idiom is needed to confirm key is not present
	// since a valid value for a key could be the zero value for the type
	fmt.Println("b in m2:", m2["b"])
}
