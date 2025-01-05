package main

import (
	"fmt"
	"io"
	"os"
)

type MyInt int

func doTypeSwitch(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Printf("i is nil. %T, %T \n", i, j)
	case int:
		fmt.Printf("i is int. %T, %T \n", i, j)
	case MyInt:
		fmt.Printf("i is MyInt. %T, %T \n", i, j)
	case io.Reader:
		fmt.Printf("i is io.Reader. %T, %T \n", i, j)
	case string:
		fmt.Printf("i is string. %T, %T \n", i, j)
	case bool, rune:
		fmt.Printf("i is bool or rune. %T, %T \n", i, j)
	default:
		fmt.Printf("i is other. %T, %T \n", i, j)
	}
}

func main() {
	var i any
	doTypeSwitch(i)

	var mine MyInt = 20
	i = mine
	doTypeSwitch(i)

	s := "This is string"
	i = s
	doTypeSwitch(i)

	s2 := []rune(s)
	doTypeSwitch(s2)

	doTypeSwitch(s2[0])

	b := int(mine) > 20
	doTypeSwitch(b)

	b2 := int(mine) == 20
	doTypeSwitch(b2)

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	st := Person{
		FirstName: "John",
		LastName:  "Baker",
		Age:       20,
	}
	doTypeSwitch(st)

	f, err := os.Open("7101.go")

	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	doTypeSwitch(f)

	finalInt := 100
	doTypeSwitch(finalInt)
}
