package main

import "fmt"

func main() {
	// Flow control
	num := 45

	// Conditional IF - ELSE-IF - ELSE
	if num < 40 {
		fmt.Println("I execute the code inside the IF")
	} else if num != 45 {
		fmt.Println("I execute the code inside the ELSE IF")
	} else {
		fmt.Println("I execute the code inside the ELSE")
	}

	age := 30

	if age >= 18 {
		fmt.Println("You are of legal age")
	} else if age < 18 && age >= 0 {
		fmt.Println("You are a minor")
	} else {
		fmt.Println("Age is invalid")
	}
}
