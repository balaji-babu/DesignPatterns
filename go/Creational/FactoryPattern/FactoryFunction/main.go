package main

import "fmt"

type Bike struct {
	Company, Model, Color string
	Wheel                 int
}

func NewBike(company, model, color string) *Bike {
	return &Bike{
		Company: company,
		Model:   model,
		Color:   color,
		Wheel:   2,
	}
}

func main() {
	b := Bike{"Bajaj", "Pulsar", "Red", 2} // Bike will have only 2 wheels, so it is very obvious to add as a default value using Factory Function(Constructor)
	fmt.Printf("Bike Details: %#v \n", b)

	nb := NewBike("RoyalEnfield", "Classic 500", "Black")
	fmt.Printf("New Bike Details: %#v \n", nb)
	// Additional support wheel might be added for a person with disabilities
	nb.Wheel = 4
	fmt.Printf("New Bike Details: %#v \n", nb)

}
