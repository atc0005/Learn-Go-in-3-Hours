package main

import "fmt"

type Foo struct {

	// due to uppercase leading character, this is "public"
	// Go encourages public fields for simple data structures
	A int
	b string
}

type Bar struct {
	C string

	// declare field of type Foo (which is itself a struct)
	F Foo
}

type Baz struct {
	D string

	// embedded struct
	Foo
	// embedding is a very useful way to reuse a struct across multiple
	// structs. Imagine defining an address struct, or a person struct, and
	// embedding it into many different types avoiding a lot of nested dotted
	// paths through struct references. While embedding might look like
	// inheritance it isn't, embedding provides delegation. Baz is not a
	// subtype of foo, it just contains a foo as a field.
	//
	// From "Introducing Go":
	//
	// Their example is a Person struct and an Android struct which embeds
	// the Person struct (aka, an "anonymous" field). This is referred to as
	// an "is a" relationship.
}

func main() {

	f := Foo{A: 10, b: "Hello"}

	// we use the initialized 'f' struct as a value for the 'F' field below
	b1 := Bar{C: "Fred", F: f}

	// we need to use the intermediate struct 'F' in order to access 'A'
	// Q: Why?
	// Q: Is this because the struct 'F' is bundled within the enclosing struct
	//    without "collapsing" the field values?
	fmt.Println(b1.F.A)

	b2 := Baz{D: "Nancy", Foo: f}

	// the intermediate struct is not needed to reference 'A' from the
	// embedded struct.
	// Q: Are the embedded struct fields "collapsed" into the enclosing struct?
	// A: According to "Introducing Go", this is an "is a" relationship with
	//    their example being a 'Person' struct and an 'Android' struct which
	//    embeds the 'Person' struct (aka, an "anonymous" field). This reflects
	//    that an Android "is a" Person, so naturally has the same fields and
	//    (presumably in later code) the same methods?
	fmt.Println(b2.A)

	// declare and initialize 'f2' struct, set values equal to b2.Foo struct
	var f2 Foo = b2.Foo
	fmt.Println(f2)

}
