package main

import "fmt"

func modMap(m map[int]string) {
	m[2] = "Hello"
	m[3] = "Goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "No. 1",
		2: "No. 2",
	}
	fmt.Print(m)

	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	fmt.Print(s)

	modSlice(s)
	fmt.Println(s)
}
