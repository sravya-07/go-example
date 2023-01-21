// Material taken from https://gobyexample.com/variadic-functions
// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

package main

import (
	"fmt"
	"reflect"
)

func sum(nums ...int) int {
	// Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
	fmt.Println(nums)
	fmt.Println(reflect.TypeOf(nums)) // []int

	sum_of_nums := 0
	for _, value := range nums { // cant use a var like index instead of _ because go will throw an error if it is not used
		sum_of_nums += value
	}

	return sum_of_nums
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr...))        // if you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	fmt.Println(sum(1, 2, 3, 4, 5)) // Variadic functions can be called in the usual way with individual arguments
}
