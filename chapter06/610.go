package main

import (
	"fmt"
)

func stringp(s string) *string {
	return &s
}

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	var y *int
	fmt.Println(y == nil)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Petterson",
	}

	fmt.Println(p)
}
