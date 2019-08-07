package main

import "fmt"

type tester interface {
	// List the method signatures that make a type meet our interface. The
	// method signatures include a method name, the parameter types, and the
	// return types.
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {

		// AND the last test result with the current one, declaring success
		// if both tests pass. Using this approach will result in a final
		// pass or fail score for the whole slice of "tester" type.
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

func (rt rangeTest) test(i int) bool {
	fmt.Printf("rt is %+v\n", rt)
	return rt.min <= i && i <= rt.max
}

type divTest int

func (dt divTest) test(i int) bool {
	fmt.Printf("Testing whether %d mod %d is 0\n", i, i)
	return i%int(dt) == 0
}

func main() {
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 20},
		divTest(5),
	})
	fmt.Printf("Final test results: %v\n", result)
}
