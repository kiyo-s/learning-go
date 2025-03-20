package adder_test

import (
	"adder"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Errorf("Invalid result: expected %d, got %d", 5, result)
	}
}
