package main

import "fmt"

func main() {
	{ // outer block
		x := 30
		fmt.Println(x) // 30
		{              // inner block
			x = 20 // overrides the value of x 30 -> 20
			y := 10
			fmt.Println(x, y) // 20, 10
		}
		fmt.Println(x) // 20
		// fmt.Println(y) // error -> undeclared name: y
	}
}
