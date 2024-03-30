package main

import "fmt"

func main() {
	var i int

	fmt.Print("Enter a value: ")
	count, err := fmt.Scanf("%d", &i)
	fmt.Println("\ncount= ", count, " err= ", err)

	switch i {
	case 100:
		fmt.Println("The number is 100")
	case 200, 300:
		fmt.Println("The number is 200 or 300")
	default:
		fmt.Println("The number is neither 100 or 200 or 300")
	}

	switch i {
	case 100:
		fmt.Println("Number is 100 or >")
		fallthrough
	case 99:
		fmt.Println("Number is 99 or >")
		fallthrough
	case 98:
		fmt.Println("Number is 98 or >")
		fallthrough
	default:
		fmt.Println("default")
	}

	var j, k int
	fmt.Println("Enter value for j & k: ")
	fmt.Scanf("%d %d", &j, &k)

	switch {
	case j+k <= 100:
		fmt.Println("<= 100") // only this condition will be satisfied and not all the other ones even though they all evaulate to true because go implicitly includes a break after each case
	case j+k <= 200:
		fmt.Println("<= 200")
	case j+k <= 50:
		fmt.Println("<=50")
	default:
		fmt.Println("Hello")
	}

}
