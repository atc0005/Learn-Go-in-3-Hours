package main

import "fmt"

// Per Learn Go in 3 Hours course:
// Go doesn't have optional parameters, named parameters or function
// overloading

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println("in doubleFail:", a, arr, s)
}
func main() {

	a := 1
	arr := [2]int{2, 4}
	s := "hello"

	doubleFail(a, arr, s)
	fmt.Println("in main:", a, arr, s)

}
