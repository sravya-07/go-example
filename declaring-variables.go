package main

import (
	"fmt"
	"reflect"
)

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

	// Finding the type of a variable
	// %T format specifier
	// reflect.TypeOf
	var (
		j int     = 10
		k string  = "Hello"
		l float64 = 1.234
		m bool    = true
	)
	fmt.Printf("j = %v, type = %T\n", j, j)
	fmt.Printf("k = %v, type = %T\n", k, k)
	fmt.Printf("l = %v, type = %T\n", l, l)
	fmt.Printf("m = %v, type = %T\n", m, m)

	var n int = 20
	var p string = "Hello"

	fmt.Println("n = %v, type = %v", n, reflect.TypeOf(n))
	fmt.Println("p = %v, type = %v", p, reflect.TypeOf(p))

	var q = "john"
	fmt.Println(q)
}
