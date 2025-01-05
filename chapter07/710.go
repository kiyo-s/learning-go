package main

import "fmt"

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s : age %d", p.LastName, p.FirstName, p.Age)
}

func main() {
	p := Person{
		LastName:  "Takeda",
		FirstName: "Shingen",
		Age:       52,
	}

	output := p.String()
	fmt.Println(output)
}
