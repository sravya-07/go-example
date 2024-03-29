// go run type-casting.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// type casting -> converting one data type to another
	// 				-> conversion doesn't mean the value will remain intact

	var (
		a int     = 90
		b float64 = float64(a) // int to float
	)

	fmt.Printf("a = %v, b = %0.6f, typeof b = %T", a, b, b)

	var (
		c float64 = 12.345
		d int     = int(c)
	)

	fmt.Printf("\nc = %v, d = %v, typeof d = %T", c, d, d)

	//strconv package to convert to & from strings
	var e int = 20
	var f string = strconv.Itoa(e)

	fmt.Printf("\nf = %q, typeof of f = %T", f, f)

	// var g int = strconv.Atoi(f)
	g, err := strconv.Atoi(f)
	fmt.Printf("\ng = %v, typeof g = %T\nerr = %v, typeof err = %T", g, g, err, err)

	h, err := strconv.Atoi("200abc")
	fmt.Printf("\nh = %v, typeof h = %T", h, h)
	fmt.Printf("\nerr = %v, typeof err = %T", err, err)
}
