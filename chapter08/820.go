package main

import (
	"errors"
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are allowed")
	}
	return i * 2, nil
}

func main() {
	i := 20
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%d doubled is %d\n", i, double)
}
