package main

import (
	"fmt"
)

func multi() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(multi())
	a, b := multi()
	fmt.Println(a, b)
	_, c := multi()
	// fmt.Println(_, c) // will fail cannot use _ as value or type
	fmt.Println(c)
}
