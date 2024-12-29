package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Printf("%s is %b\n", string(b), b)

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]

	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	var str string = "Hello ☀"
	var str2 string = str[4:7]
	var str3 string = str[:5]
	var str4 string = str[6:]

	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
}
