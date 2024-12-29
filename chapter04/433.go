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

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Loop: ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}
