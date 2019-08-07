package main

import "fmt"

func doStuff(i interface{}) {

	// A type switch looks a great deal at this switch statement that we saw
	// way back in section 3. For a type switch we have the keyword switch,
	// and we do a type assertion like we saw in the previous example, only
	// instead of using the name of a concrete type you put the keyword type
	// itself in the parenthesis instead.
	//
	// So by convention, the variable on the left-hand side of a colon equals
	// reuses the interface variables name on the right-hand side. And I agree
	// this convention is a little bit confusing, but you get used to it after
	// a bit.
	switch i := i.(type) {

	// each case statement has a type name on it
	case int:
		fmt.Println("Double i is", i+i)
	case string:
		fmt.Println("i is", len(i), "characters long")

	// default indicates that we do not know what type is behind the interface
	default:
		fmt.Println("I don't know what to do with this")
	}
}

func main() {
	doStuff(10)
	doStuff("Hello")
	doStuff(true)
}
