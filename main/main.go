package main

import "fmt"

func main() {
	myFunction()
	greeting("Inés")
	fmt.Println("The addition is: ", addition(5, 7))
}

// Function that does not receive a value
func myFunction() {
	fmt.Println("I write from my function")
}

// Function that receives a value
func greeting(name string) {
	fmt.Println("Hello.... " + name)
}

// Function that returns a value
func addition(num1, num2 int) int {
	return num1 + num2
}
