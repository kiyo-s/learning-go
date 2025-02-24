package main

import (
	"fmt"

	print "example.co.jp/package_example/formatter"
	"example.co.jp/package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
