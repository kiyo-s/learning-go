package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		v := 1
		fmt.Printf("will import %v to ch1\n", v)
		ch1 <- v
	}()

	go func() {
		v := 2
		fmt.Printf("will import %v to ch2\n", v)
		ch2 <- v
	}()

	go func() {
		fmt.Printf("weill receive value from ch3\n")
		v := <-ch3
		fmt.Printf("received %v from ch3\n", v)
	}()

	go func() {
		v := 4
		fmt.Printf("will send %v to ch4\n", v)
		ch4 <- v
		fmt.Printf("sended %v to ch4\n", v)
	}()

	x := 3

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	case ch3 <- x:
		fmt.Println("send to ch3:", x)
	case <-ch4:
		fmt.Println("received from ch4, but ignored")
	}
}
