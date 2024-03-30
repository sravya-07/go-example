package main

import "fmt"

func main() {
	var num = 7

	if num%2 == 0 { // we can use if blocks independently regardless of the other blocks
		fmt.Println("Even number: ", num)
	} else if num%3 == 0 { // else & else ifshould always start where the last bracket of if ends
		fmt.Println("Odd number & it's divisible by 3: ", num)
	} else {
		fmt.Println("Odd number & it's not divisible by 3: ", num)
	}

	var a, b string = "foo", "bar"
	if a+b == "foo" {
		fmt.Println("foo")
	} else if a+b == "foobar" { // when the first if else matches it will not check the rest of them the flow of execution leaves the if elseif else nest and proceeds to the next statement
		fmt.Println("bar")
	} else if a+b == "foobar" {
		fmt.Println("foobar")
	} else {
		fmt.Println("None matched")
	}
	fmt.Println("thank you!")
}
