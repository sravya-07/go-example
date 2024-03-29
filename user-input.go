package main

import "fmt"

func main() {
	// fmt.Scanf("%<format specifier> (s)", object_arguments)
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name) // pointer used

	// fmt.Println("Your name is ", name)

	// var (
	// 	name1   string
	// 	has_car bool
	// )

	// fmt.Println("Enter you name and whether you have a car or not ?")
	// fmt.Scanf("%s %t", &name1, &has_car) // reads from input source sequentially

	// fmt.Println(name1, "has car: ", has_car)

	/*
		Scanf -> return 2 values
			count -> number of variables that it writes to
			err -> any error during execution
	*/

	var (
		a string
		b int
	)

	fmt.Println("Enter a string & a num: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("Count: ", count, "\nerror: ", err, "\na = ", a, "\nb = ", b)
}
