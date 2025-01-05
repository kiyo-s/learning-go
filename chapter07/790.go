package main

import "fmt"

func main() {
	var i any
	i = 20
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	i = struct {
		FirtstName string
		LastName   string
	}{"Shingen", "Takeda"}
	fmt.Println(i)
}
