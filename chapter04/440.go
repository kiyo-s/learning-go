package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain",
		"タコの足とイカの足", "antholopologist", "タコの足は8本でイカの足は10本"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("\"%s\" has %d characters, too short.\n", word, size)
		case 5:
			fmt.Printf("\"%s\" has %d characters, just right.\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("\"%s\" has %d characters, too long.\n", word, size)
			if n := len(word); size < n {
				fmt.Printf("The word is %d bytes long.\n", n)
			} else {
				fmt.Println()
			}
		}
	}
}
