package main

import (
	"fmt"
)

func main() {
	sample_map := make(map[string]int) // string -> key type int -> value type
	// make(map[key-type]value-type)
	sample_map["one"] = 1
	sample_map["two"] = 2

	fmt.Println(sample_map)

	new_map := map[string]string{"name": "Sravya"}
	fmt.Print(new_map)
}
