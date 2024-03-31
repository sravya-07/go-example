package main

import "fmt"

func main() {
	// slices
	// slice is quite literally a slice of a larger array
	// a slice has a pointer (pointed to the element in the array from where the slice first starts)
	//			   len & capacity
	//             len -> the length of the slice
	//             capacity -> how many elements there are in the underlying array if the starting point is the first elem of the slice

	// menthod 1 of slice init
	grades := []int{10, 20, 30}
	var grades1 []int = []int{10, 20, 30, 40}

	fmt.Print("grades= ", grades)
	fmt.Printf(" type=%T", grades)
	fmt.Print("\ngrades1= ", grades1)
	fmt.Printf(" type=%T", grades1)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\narr = ", arr)
	slice_1 := arr[1:8]
	fmt.Println("slice_1 ", slice_1)

	sub_slice_1 := slice_1[0:3]
	fmt.Println("sub slice 1 ", sub_slice_1)

	// slice := make([]<data_type>, length, capacity)
	slice_2 := make([]int, 5, 10)
	fmt.Print("slice2= ", slice_2)
	fmt.Printf(" len=%d cap%d\n", len(slice_2), cap(slice_2))

	arr3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 := arr3[1:8]
	fmt.Println(cap(slice3))
	fmt.Println(cap(arr3))

	// value change
	slice3[4] = 3489759845
	fmt.Println(slice3, "\n", arr3)

	slice4 := append(slice3, 45)
	fmt.Println("arr3=", arr3)
	fmt.Println("slice4=", slice4)

	slice5 := append(slice3, 1, 1, 1, 1)
	fmt.Println("arr3=", arr3)
	fmt.Println("slice3=", slice3)
	fmt.Println("slice5=", slice5)
}
