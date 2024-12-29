package main

import "fmt"

func main() {
	var nilMap map[string]int

	fmt.Println(nilMap)
	fmt.Println(nilMap == nil)
	fmt.Println(len(nilMap))

	totalWins := map[string]int{}
	fmt.Println(totalWins)
	fmt.Println(totalWins == nil)
	fmt.Println(len(totalWins))
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"])

	teams := map[string][]string{
		"Writers": []string{"Natsume", "Mori", "Kunikiad"},
		"Knigits": []string{"Takeda", "Tokugawa", "Akechi"},
		"Musician": []string{"Mozart", "Beethoven", "Bach"},
	}
	fmt.Println(teams["Writers"][0])

	ages := make(map[string]int, 10)
	fmt.Println(len(ages))

	totalWins["Writers"] = 1
	totalWins["Knigits"] = 2
	fmt.Println("Writers: ", totalWins["Writers"])
	fmt.Println("Knigits: ", totalWins["Knigits"])
	fmt.Println("Musicians: ", totalWins["Musicians"])

	totalWins["Musicians"]++
	fmt.Println("Writers: ", totalWins["Writers"])
	fmt.Println("Knigits: ", totalWins["Knigits"])
	fmt.Println("Musicians: ", totalWins["Musicians"])

	totalWins["Knigits"] = 3
	fmt.Println("Writers: ", totalWins["Writers"])
	fmt.Println("Knigits: ", totalWins["Knigits"])
	fmt.Println("Musicians: ", totalWins["Musicians"])

	// Comma ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	delete(m, "hello")

	v, ok = m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	// set
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
		fmt.Println(v)
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
}
