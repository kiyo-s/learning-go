package main

import (
	"fmt"
	"reflect"
)

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Adviertisements
)

func main() {
	var m MailCategory = Personal
	fmt.Println("Personal:", m)
	m = Uncategorized
	fmt.Println("Uncategorized:", m)
	fmt.Println(reflect.TypeOf(m))
}
