package main

import (
	"fmt"
)

func main() {

	//Mathematical operators
	addition := 10 + 5
	subtraction := 10 - 5
	multiplication := 10 * 5
	division := 10 / 5
	module := 10 % 5
	fmt.Println(addition, subtraction, multiplication, division, module)

	// Increment operators
	addition++
	fmt.Println(addition)
	addition--
	fmt.Println(addition)

	// Assignment operator with addition
	addition += 10
	fmt.Println(addition)

	// Comparison operators
	fmt.Println(10 > 5)   // true
	fmt.Println(10 < 5)   // false
	fmt.Println(10 >= 5)  // true
	fmt.Println(10 <= 5)  // false
	fmt.Println(10 == 5)  // false
	fmt.Println(10 != 5)  // true
	fmt.Println(10 != 10) // false

	// Logical operators
	fmt.Println(10 < 5 && 10 > 4) // false
	fmt.Println(10 > 5 && 10 > 4) // true
	fmt.Println(10 < 5 || 10 > 4) // true
}
