package main

import "fmt"

type myInt int

func (mi myInt) isEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}

func main() {

	// type conversion of literal 10 to myInt type; we can call our
	// methods against the myInt type (with a value of 10)
	m := myInt(10)

	fmt.Println(m)
	fmt.Println(m.isEven())
	m.Double()
	fmt.Println(m)

}
