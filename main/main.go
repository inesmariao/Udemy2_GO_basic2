package main

import "fmt"

func main() {
	/*
		Programa que recorra los números del 0 al 20 y a través de un mensaje por consola muestre el número y muestre solo los que son pares.

		Program that goes through the numbers from 0 to 20 and through a console message displays the number and shows only those that are even.
	*/

	fmt.Println("Números pares del 0 al 20: ")
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d es un número par.\n", i)
		}
	}
}
