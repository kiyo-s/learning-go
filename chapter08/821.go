package main

import (
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("only even numbers are allowed, %d is not even", i)
	}
	return i * 2, nil
}

func main() {
	i := 21
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%d doubled is %d\n", i, double)
}
