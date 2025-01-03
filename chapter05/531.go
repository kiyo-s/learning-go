package main

import "fmt"

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i <= 5; i++ {
		fmt.Printf("%d: %d, %d\n", i, twoBase(i), threeBase(i))
	}
}
