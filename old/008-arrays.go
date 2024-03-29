package main

import (
	"fmt"
)

func main() {
	var a[5]int
	fmt.Print(a) // initialized to all zeros

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	// get length
	fmt.Println(len(a))

	// var b [3]int [1,2,3]
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [2][2]int{{1,2},{3,4}}
	fmt.Println(c)
}