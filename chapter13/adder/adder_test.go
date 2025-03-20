package adder

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Preparing for tests")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Cleaning up after tests")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses what was set in TestMain:", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond also uses what was set in TestMain:", testTime)
}

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Errorf("Invalid result: expected %d, got %d", 5, result)
	}
}
