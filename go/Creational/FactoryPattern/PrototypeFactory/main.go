package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 500000}
	case Manager:
		return &Employee{"", "manager", 8000000}
	default:
		panic("Unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Jack"
	fmt.Printf("Manager Details: %#v\n", m)
}
