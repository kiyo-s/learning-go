package bar

import "github.com/learning-go/internal_example/foo/internal"

func FailureUseDouble(i int) int {
	return internal.Doubler(i)
}
