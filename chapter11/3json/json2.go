package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	data := readData()
	var o Order
	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", o)

	out, err := json.Marshal(o)
	fmt.Println(out)
}

func readData() []byte {
	b, err := os.ReadFile("testdata/data.json")
	if err != nil {
		log.Fatal(err)
	}
	return b
}
