package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	funcs := prepareFunctions()
	s := "Time flies like an arrow."
	fmt.Println("Original string is: ", s)
	r := convertData(s, funcs)
	fmt.Printf("First result is: %v\n\n", r)

	fmt.Println("Wait for end of other goroutine.")
	time.Sleep(1 * time.Second)

	fmt.Println("main function end.")
}

type message struct {
	s        string
	fromFunc int
}

func convertData(s string, converters []func(string) message) message {
	done := make(chan struct{})
	resultChan := make(chan message)

	for _, f := range converters {
		go func(f func(string) message) {
			r := f(s)
			select {
			case resultChan <- r:
				fmt.Printf("r is %v\n", r)
			case <-done:
				fmt.Println("case <-done select ", r.fromFunc)
			}
			fmt.Println("Non-named function end.")
		}(f)
	}

	r := <-resultChan
	close(done)
	return r
}

func randomPeriod() time.Duration {
	t := time.Millisecond * time.Duration(rand.Intn(4)+2)
	return t
}

func prepareFunctions() []func(string) message {
	funcNo1 := func(a string) message {
		sleepRandomPeriod(1)
		b := strings.ToLower(a)    // 小文字にする
		fmt.Println("処理完了（1）:", b) // 1: time flies like an arrow.
		return message{b, 1}
	}
	funcNo2 := func(a string) message { // 大文字にする
		sleepRandomPeriod(2)
		b := strings.ToUpper(a)
		fmt.Println("処理完了（2）:", b) // 2: TIME FLIES LIKE AN ARROW.
		return message{b, 2}
	}
	funcNo3 := func(a string) message { // i -> I
		sleepRandomPeriod(3)
		b := strings.ReplaceAll(a, "i", "I")
		fmt.Println("処理完了（3）:", b) // 3: TIme flIes lIke an arrow.
		return message{b, 3}
	}
	funcNo4 := func(a string) message { // e->E  r->L
		sleepRandomPeriod(4)
		b := strings.ReplaceAll(a, "e", "E")
		a = strings.ReplaceAll(a, "r", "L")
		fmt.Println("処理完了（4）:", b) // 4:
		return message{b, 4}
	}

	return []func(string) message{funcNo1, funcNo2, funcNo3, funcNo4}
}

func sleepRandomPeriod(funcNum int) {
	t := randomPeriod()
	fmt.Println(funcNum, "will sleep: ", t)
	time.Sleep(t)
}
