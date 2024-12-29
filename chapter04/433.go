package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Hanako": true, "Taro": true, "Yoko": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}
