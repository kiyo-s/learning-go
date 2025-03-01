package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer func() {
			close(ch)
			fmt.Println("Closed ch")
		}()

		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Write to ch ", i)
		}
	}()

	for v := range ch {
		fmt.Println("Receive from ch ", v)
	}
}
