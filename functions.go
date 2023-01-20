package main

import (
	"fmt"
)

// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
func add(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func addadd(a, b, c int) int {
	return a + b + c
}

func main() {
	var a, b int
	a = 2
	b = 3
	fmt.Println(a, b, add(a, b), addadd(a, b, b))
}
