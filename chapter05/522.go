package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Printf("%d in non-named function\n", j)
		}(i)
	}
}
