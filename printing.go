package main

import "fmt"

func main() {
	fmt.Print(1)
	fmt.Print("Hello")

	var s1 string = "Hello World"
	fmt.Print("\n")
	fmt.Print(s1)

	var i int = 10
	fmt.Print("\n", i)

	var j int = 20
	fmt.Println(j) // adds a newline after printing
	var k int = 30
	fmt.Print(k)

	// fmt.Printf("template string %s", object args(s))
	/*
		%v -> formats value in default format
		%d -> formats decimal integers
		%T -> Type of value
		%c -> Character
		%q -> Quoted characters/string
		%s -> plain string
		%f -> floating numebers
		%0.2f -> floating numbers upto 2 decimal places
		%t -> bool
	*/
	fmt.Printf("Hi %v", "hello")
	fmt.Printf("Hi %d", 10)
	fmt.Printf("Hi %v your marks are %d", "Sravya", 10) // write the variables in the order of where you want them to appear in the string
}
