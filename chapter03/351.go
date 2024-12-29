package main

import "fmt"

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Bob"
	person.age = 15
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "Pochi",
		kind: "dog",
	}

	fmt.Println(person)
	fmt.Println(pet)
}
