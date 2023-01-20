package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)

	for index, item := range nums {
		fmt.Println("index:", index, "item:", item)
	}

	mapping := map[int]string{1: "monday", 2: "tuesday"}
	for key, value := range mapping {
		fmt.Println("key:", key, "value:", value)
	}
}
