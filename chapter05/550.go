package main

import "fmt"

type Person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p Person) {
	i *= 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	p := Person{}
	i := 2
	s := "Hello"
	fmt.Println(i, s, p)
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}
