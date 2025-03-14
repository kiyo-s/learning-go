package main

import (
	"fmt"
	"time"
)

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func main() {
	for i := range countTo(10) {
		fmt.Print(i, " ")
	}
	fmt.Println()
	doSomethingTakingLongTime()
}

func doSomethingTakingLongTime() {
	time.Sleep(5 * time.Second)
}
