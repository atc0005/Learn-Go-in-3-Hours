package main

import "fmt"

type Foo struct {
	A int
	B string
}

func (f Foo) fieldCount() int {
	fmt.Println("    Inside Foo fieldCount() method")
	return 2
}

func (f Foo) String() string {
	fmt.Println("    Inside Foo String() method")
	return fmt.Sprintf("%d fields: A: %d and B: %s", f.fieldCount(), f.A, f.B)
}

func (f *Foo) Double() {
	fmt.Println("    Inside Foo Double() method")
	f.A = f.A * 2
}

type Bar struct {
	C bool

	// embedded struct of type 'Foo'
	Foo
}

func (b Bar) String() string {

	// b.Foo.String() is a call to the 'String()' method on the embedded 'Foo'
	//
	// Q: I thought this was not needed/desired for embedded structs? e.g.,
	//    the "is a" relationship?
	// !: Note: If you try to call 'b.String()' here it results in a 'stack
	//    overflow' error, presumably due to infinite recursion
	//return fmt.Sprintf("%s and C: %v", b.Foo.String(), b.C)
	fmt.Println("    Inside Bar String() method")
	return fmt.Sprintf("%s and C: %v", b.Foo.String(), b.C)
}

// Note: Never going to be called from inside Foo; later when the
// f.fieldCount() method is called it is called against Foo and not Bar; Foo
// has no idea that Bar has a fieldCount() method, even with Foo embedded
// inside of it. See notes below in main()
//
// This is confusing for users of languages where overriding is supported.
func (b Bar) fieldCount() int {
	fmt.Println("    Inside Bar fieldCount() method")
	return 3
}

func main() {

	f := Foo{
		A: 10,
		B: "Hello",
	}

	b := Bar{
		C:   true,
		Foo: f,
	}

	// This calls the String() method on 'b', which in turn explicitly calls
	// the String() method on embedded 'f'.
	//
	// The String() method on embedded 'f' calls fieldCount() on method
	// receiver 'f' instead of the fieldCount() method on 'b'. There is no
	// inheritance in Go.
	// Foo has no idea that there's a fieldCount() method on Bar as well even
	// though Foo is embedded inside of it. That's because in Go there's no
	// method of overriding and there's no virtual method dispatch. So the
	// fieldCount() method that we declared on Bar is never going to be called
	// from inside of Foo.
	fmt.Println(b.String())

	// embedded access on Foo struct
	// ... if there's a method declared on an embedded struct, we are able
	// to access the method on the outer struct as though the method was
	// declared on the outer struct.
	//
	// Note: Presumably this is due to the "is a" relationship (e.g., an
	// "Android" is a "Person" with the same base traits)
	b.Double()

	// Same as earlier example
	fmt.Println(b.String())

}
