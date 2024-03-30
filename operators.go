package main

import "fmt"

// import "fmt"

func main() {
	// operators -> help in performing specific logical, mathematical operations

	// a + b a, b operands -> on whom the operation is performed
	// types of operators
	// comparison operators == != < <= > >=
	// arithmetic operators + - * / % ++ --
	// assignment operators = += -= *= /= %=

	// comparion operators
	// compare 2 values and yeild a boolean value true/false

	var a string = "Hello"
	var b string = "skfvhs"
	fmt.Println(a == b) // false

	var c = "Hello"
	fmt.Println(a == c) // true

	fmt.Println(a != b) // true

	fmt.Println(5 < 10)  // true
	fmt.Println(5 <= 10) // true

	fmt.Println(10 > 5)   // true
	fmt.Println(10 >= 5)  // true
	fmt.Println(10 >= 10) // true

	// fmt.Println("a" > 10) error: invalid operation: cannot compare "a" > 10 (mismatched types untyped string and untyped int)

	// arithmetic operators + - * / %
	// these operations can only be performed on values that are of the same datatype, if they are different then it wont be valid
	// + addition
	var d = "Hello"
	var e string = " World"

	fmt.Println(d + e) // + operator concatenates 2 strings in go lang
	fmt.Println(2 + 3)

	// -
	// fmt.Println(d - e) -> error: invalid operation: operator - not defined on d (variable of type string)
	fmt.Println(4 - 5)

	// * % /
	fmt.Println(2 * 3)
	fmt.Println(12 / 2)
	fmt.Println(15 / 2)     // integer division -> integer output
	fmt.Println(15.0 / 2.0) // float division float output
	fmt.Println(13 % 2)
	// fmt.Println(13.0 % 2.0) -> invalid operation: operator % not defined on 13.0 (untyped float constant 13)

	// unary operators -> work on single variable to change the value
	var (
		f = 11
		g = 9
	)
	f--
	g++
	fmt.Println(f, g)
	var h = 1.2
	h++
	fmt.Println(h) // 2.2

	// logical operators
	// && || !
	var i = 10
	fmt.Println((i < 100) && (i < 99))
	fmt.Println((i < 100) && (i < 1))
	fmt.Println(!((i < 100) && (i < 1)))
	fmt.Println((i < 100) || (i < 1))

	// assignement operators
	// = += -= *= /= %=
	var j = 10
	var k int
	k = j

	k = k + j
	k += j

	k = k - j
	k -= j

	k = k * j
	k *= j

	k = k / j
	k /= j // assign the quotient

	k = k % j
	k %= j

	// bitwise operators
	// & | ^ >> <<
	// & -> two inputs and does an AND operation on every bit of the two numbers
	// 12 = 00001100 (convert to binary)
	// 25 = 00011001
	//   00001100
	// & 00011001
	//   00001000 (logic: (0 and 0 = 0) ( 1 and 0 = 0) (1 and 1 = 1) )
	//   8 (convert to decimal)

	// | -> performs an OR operation on every bit of the two numbers
	// 12 = 00001100 (convert to binary)
	// 25 = 00011001
	//   00001100
	// | 00011001
	//   00011101
	//   29

	// ^ -> bitwise xor operation
	// takes 2 numbers and does XOR operation on the two numbers
	// XOR = 1 if the two bits are opposite otherwise 1
	//   00001100
	// ^ 00011001
	//   00010101
	//   21 in decimal

	// << left shift operator
	// shifts all the bits left by a speccified number of bits
	// the bit positions that have been vacated by the the left shift operator are filled with 0
	// 212 << 1
	//    11010100 (binary)
	//   110101000 (extra zero added in the position vacated by the left shift operator)
	//   424 (decimal)

	// >> right shift operator
	// shifts all the bits to right by a specified number of bits
	// 212 >> 2
	//    11010100 (binary)
	//     00110101 (excess bitsa shifted off to the right are discarded)
	//     53 in decimal

	fmt.Println(25 & 12)
	fmt.Println(12 & 25)
	fmt.Println(12 | 25)
	fmt.Println(25 | 12)
	fmt.Println(12 ^ 25)
	fmt.Println(212 << 1)
	fmt.Println(212 >> 2)
}
