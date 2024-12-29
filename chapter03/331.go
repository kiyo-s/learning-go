package main

import "fmt"

func main() {
	// rune
	var a rune = 'x'
	var s string = string(a)

	var b byte = 'y'
	var s2 string = string(b)

	var i int = 65
	var s3 string = string(i)

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

	var str = "Hello, 🌞"
	var bs []byte = []byte(str)
	var rs []rune = []rune(str)
	fmt.Println(bs)
	fmt.Println(rs)
}
