package main

import "fmt"

func main() {
	// declaring a variable ;
	var a int
	a = 10
	fmt.Println(a)

	// declaring & initialzing
	var b int = 20
	fmt.Println(b)

	var c string
	// c = 30; error -> cannot use value 30 as type string in assignment
	c = "Hello"
	fmt.Println(c)

	// short variable declaration
	s := "Hello World"
	fmt.Println(s)
	fmt.Printf("%s %T\n", s, s)

	// d = "Hi Hello" -> error undeclared variable d
	d := "Hi Hello"
	d = "New value"
	fmt.Println(d)

	// shorthand way of declaring multiple vars of the same type
	var e, f string = "foo", "bar"
	fmt.Println(e, f)

	var g int = 40
	// g = "Hello" error -> cannot use "Hello" (untyped string constant) as int value in assignment
	fmt.Println(g)

	// var b bool = "false" -> false is in double quotes its wrong  var b bool = "false"

	var (
		h int    = 10
		i string = "Hello"
	)
	fmt.Println(h, i)
}
