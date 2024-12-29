package main

import "fmt"

func main() {
	for l := 1; l <= 100; l++ {
		if l%3 == 0 {
			if l%5 == 0 {
				fmt.Println("FizzBUzz: ", l)
			} else {
				fmt.Println("Fizz: ", l)
			}
		} else if l%5 == 0 {
			fmt.Println("Buzz: ", l)
		} else {
			fmt.Println(l)
		}
	}
}
