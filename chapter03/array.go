package main

import "fmt"

func main(){
    var x [3]int
    fmt.Println(x)
    fmt.Println(len(x))

	var y = [3]int{10, 20, 30}
	fmt.Println(y)
    fmt.Println(len(y))

	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z)
    fmt.Println(len(z))

	var a = [...]int{10, 20, 30}
	var b = [3]int{10, 20, 30}
	fmt.Println(a == b)

	var c [2][3]int
	fmt.Println(c)
}
