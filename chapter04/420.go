package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("too small: ", n)
	} else if n > 5 {
		fmt.Println("too big: ", n)
	} else {
		fmt.Println("just right: ", n)
	}
}
