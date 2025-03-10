package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const data = `
		{"name": "Fred", "age": 40}
		{"name": "Meary", "age": 21}
		{"name": "Bad", "age": 30}
		`
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)

	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("-----")
	out := b.String()
	fmt.Println(out)
}
