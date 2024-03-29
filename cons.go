package main

import "fmt"

const PI float64 = 3.14

func main() {
	// fmt.Println("Hello")
	// values that do not change
	const name string = "HELLO"

	// two types of constants typed & untyped

	// untyped constant -> when you do not explcitly provide a type, the type is infered
	const i = 10

	// typed constant -> when the type is explicity provided
	const j int = 10

	// const k string;
	// k = "hello"
	// not allowed error: missing constant value
	// concept of zero values doesnt apply here

	// i = i + 10
	// error: cannot assign to i (untyped int constant 10)

	// cant use := shorthand for constants

	var k = "john"
	fmt.Println(k)

}
