package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Too small: ", n)
	} else if n > 5 {
		fmt.Println("Too big: ", n)
	} else {
		fmt.Println("Just right: ", n)
	}

	// fmt.Println(n) // Error: undefined: n
}
