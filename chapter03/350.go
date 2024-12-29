package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	julia := person{
		"Julia", 40, "cat",
	}

	beth := person{
		age:  30,
		name: "Beth",
	}

	bob := person{}
	bob.name = "Bob"

	fmt.Println(julia)
	fmt.Println(beth)
	fmt.Println(bob)
}
