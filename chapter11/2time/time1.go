package main

import (
	"fmt"
	"time"
)

func main() {
	d := 2*time.Hour + 30*time.Minute + 45*time.Second
	fmt.Println(d)

	s := "3h31m46s"
	st, _ := time.ParseDuration(s)
	fmt.Println(st)
}
