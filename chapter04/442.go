package main

import "fmt"

func main() {
loop:
	for i := 1; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, " is even.")
		case i%3 == 0:
			fmt.Println(i, " is divisible by 3, but not even.")
		case i%7 == 0:
			fmt.Println(i, " is divisible by 7, but not even or divisible by 3.")
			fmt.Println("Breaking the loop.")
			break loop
		default:
			fmt.Println(i, " is boring.")
		}
	}
}
