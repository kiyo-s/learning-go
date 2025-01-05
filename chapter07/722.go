package main

import "fmt"

type Score int
type HighScore Score

func (s Score) Double() Score {
	return s * 2
}

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

type Employee Person

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	fmt.Println(s, s.Double())
	fmt.Println(Score(hs), Score(hs).Double())

	s = Score(i)
	hs = HighScore(s)

	fmt.Println(s, hs)
	hhs := hs + 20
	fmt.Println(hhs)
}
