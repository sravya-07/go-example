package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3} // {} used not []
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	var e [5]int = [5]int{}
	f := [5]int{}
	fmt.Println(e, f) // arrays filled with 0

	fmt.Println(a, b, c, len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

	for i, elem := range a { // range returns the index & the element
		fmt.Println(i, elem)
	}

	d := [3][2]int{{1, 2}, {3, 4}, {5, 6}} // 3 rows & 2 columns
	fmt.Println(d)

	for i := 0; i < len(d); i++ {
		// fmt.Println("Row ", i + 1)
		for j := 0; j < len(d[i]); j++ {
			fmt.Print(d[i][j], " ")
		}
		fmt.Println("")
	}
}
