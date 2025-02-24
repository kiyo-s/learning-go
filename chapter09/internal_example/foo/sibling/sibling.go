package sibling

import "github.com/learning-go/internal_example/foo/internal"

func AlsoUseDouble(i int) int {
	return internal.Doubler(i)
}
