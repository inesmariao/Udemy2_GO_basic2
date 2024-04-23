package main

import "fmt"

// POO struct
type person struct {
	name     string
	lastName string
	age      int
}

// Person structure methods
func (p person) greeting(msg string) {
	fmt.Println(msg + ", " + p.name)
}

func (p person) birthday() int {
	return p.age + 1
}

func main() {
	person1 := person{"Graciela", "Hernandez", 47}
	person2 := person{"Pedro", "Oliveros", 46}
	person3 := person{"Lina", "Oliveros", 40}

	fmt.Println("Person1: ", person1)
	fmt.Println("Person2: ", person2)
	fmt.Println("Person3: ", person3)

	person2.name = "Catalina"
	fmt.Println("Person1: ", person1)

	person1.greeting("!Hello!")
	person2.greeting("!Hello!")
	person3.greeting("!Hello!")

	fmt.Printf("%s cumplió años ayer, hoy tiene %d años", person1.name, person1.birthday())
}
