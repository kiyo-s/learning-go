package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Denis",
		Age:  37,
	}

	result := CreatePerson("Denis", 37)
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}

	if result.DateAdded.IsZero() {
		t.Error("DateAdded was't assigned")
	}
}
