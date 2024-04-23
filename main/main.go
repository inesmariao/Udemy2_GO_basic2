package main

import (
	"fmt"
	"strings"
)

func main() {

	fruits := []string{"Mango", "Strawberry", "Pear", "Apple", "Kiwi"}

	fmt.Println(fruits[1])

	fruits = append(fruits, "Watermelon", "Melon")

	printArray(fruits)

	fruits = append(fruits, "Pineapple")

	printArray(fruits)

	searchElement(fruits, "melon")

}

func printArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println()
}

func searchElement(arr []string, f string) {
	f = strings.Title(f)
	found := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == f {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("%s is in the list\n", f)
	} else {
		fmt.Printf("%s is not in the list\n", f)
	}
}
