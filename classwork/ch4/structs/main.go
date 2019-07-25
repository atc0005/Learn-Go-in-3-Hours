package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {

	// `json:"name"` is a struct tag and is primarily used to support
	// marshaling and unmarshaling data in and out of structs.
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	bob := `{ "name": "Bob", "age": 30 }`
	var b Person

	// json.Unmarshal and json.Marshal work with slices of bytes
	json.Unmarshal([]byte(bob), &b)
	fmt.Println("JSON data unmarshaled into struct:", b)

	bob2, _ := json.Marshal(b)
	fmt.Println("struct data marshaled into JSON:", string(bob2))
}
