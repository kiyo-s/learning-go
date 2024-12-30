package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return result, remainder, errors.New("\"CAN NOT DIVIDE BY ZERO\"")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
