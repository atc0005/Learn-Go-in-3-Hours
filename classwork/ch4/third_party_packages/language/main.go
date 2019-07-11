package main

import (
	"fmt"

	// local copy of package instead of using GOPATH as would be best practice
	"./mapper"
)

func main() {

	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))

}
