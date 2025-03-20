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
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
