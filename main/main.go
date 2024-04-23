package main

import "fmt"

func main() {

	// Flow control FOR
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	name := "Ines"

	for i := 0; i < 4; i++ {
		fmt.Println(i, string(name[i]))
	}

	// For loop used like the While loop, in Go there is no While loop
	num := 1
	for num != 0 {
		fmt.Println("num different from 0")
		fmt.Print("Enter another number: ")
		fmt.Scanln(&num)
	}
}
