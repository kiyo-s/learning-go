package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	newEmployees := []Employee{
		Employee{
			"Mitsunari Ishida",
			"13112",
		},
		Employee{
			"Ieyasu Tokugawa",
			"13115",
		},
	}
	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Hideyoshi Tokugawa",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	var eOK Employee = m.Employee
	fmt.Println(eOK)

	// var eNG Employee = m
}
