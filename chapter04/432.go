package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz: ", i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz: ", i)
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz: ", i)
			continue
		}
		fmt.Println(i)
	}
}
