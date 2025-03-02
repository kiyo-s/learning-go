package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct {
		Name string
		Age  int
	}{}

	err := json.Unmarshal([]byte(`{"name":"Onono Komachi","occurpation":"poets","age": 20}`), &f)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Nice to meet you, my name is", f.Name, ".", "I am", f.Age, "years old.")
	fmt.Printf("%+v", f)
}
