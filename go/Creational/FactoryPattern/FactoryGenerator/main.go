package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Functional Approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// Structural Approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 5000000)
	managerFactory := NewEmployeeFactory("manager", 7500000)
	developer := developerFactory("Harish")
	manager := managerFactory("Sagar soni")

	fmt.Printf("Developer Details: %#v\n", *developer)
	fmt.Printf("Manager Details: %#v\n", *manager)

	ceoFactory := NewEmployeeFactory2("CEO", 10000000)
	ceo := ceoFactory.Create("John Mathew")
	fmt.Printf("CEO Details: %#v\n", *ceo)
}
