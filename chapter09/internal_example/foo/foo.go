package foo

import "github.com/learning-go/internal_example/foo/internal"

func UseDouble(i int) int {
	return internal.Doubler(i)
}
