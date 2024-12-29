package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}

	k := 1
	for {
		fmt.Println("Hello, World!: ", k)
		k++
	}
}
