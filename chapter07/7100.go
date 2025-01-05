package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2)

	/*
		i3 := i.(string)
		fmt.Println(i3)

		i4 := i.(int)
		fmt.Println(i4)
	*/

	i4, ok := i.(int)
	if !ok {
		err := fmt.Errorf("Type of \"i\" (value is %v) is unexpected.", i)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(i4)
}
