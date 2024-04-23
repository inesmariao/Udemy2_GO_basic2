package main

import "fmt"

func main() {
	var intVar int
	var decimalVar float64
	var textVar string
	var booleanVar bool

	intVar = 5
	decimalVar = 10.4
	textVar = "Inés"
	booleanVar = true

	fmt.Println(intVar, decimalVar, textVar, booleanVar)

	// Declare and assign variable in the same instruction
	numVar := 10
	fmt.Println(numVar)

}
