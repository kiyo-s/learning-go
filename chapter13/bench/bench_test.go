package bench

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	makeData()
	exitVal := m.Run()
	os.Remove("testdata/data.txt")
	os.Exit(exitVal)
}

func makeData() {
	file, err := os.Create("testdata/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rand.Seed(1)
	for i := 0; i < 10000; i++ {
		data := makeWord(rand.Intn(10) + 1)
		file.Write(data)
	}
}

func makeWord(l int) []byte {
	out := make([]byte, l+1)
	for i := 0; i < l; i++ {
		out[i] = 'a' + byte(rand.Intn(26))
	}
	out[l] = '\n'
	return out
}

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 65204 {
		t.Error("Expect result is 65204, result is :", result)
	}
}

var blackhole int

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1_000, 10_000, 65_204, 100_000, 1_000_000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
