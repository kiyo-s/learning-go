package internal_package_example

import "github.com/learning-go/internal_example/foo/internal"

func FailUseDouble(i int) int {
	return internal.Doubler(i)
}
