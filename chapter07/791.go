package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	getData()
}

func getData() error {
	data := map[string]any{}

	contents, err := os.ReadFile("sample.json")
	// contents, err := ioutil.ReadFile("sample.json") //こちらでも可だが上が呼ばれる
	if err != nil {
		return err
	}

	json.Unmarshal(contents, &data)
	fmt.Println(data)
	return nil
}
