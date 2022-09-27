package main

import "fmt"

type Person interface {
	Booked()
}

type person struct {
	Name, Address, Model string
	MobileNo             uint
}

func (p *person) Booked() {
	fmt.Printf("Hi %s, Thanks for successfully booking your bike %s from Royal Enfield. \nYour product will be delivered in three months.", p.Name, p.Model)
}

func NewBike(name, address, model string, mobile uint) Person {
	return &person{
		Name:     name,
		Address:  address,
		Model:    model,
		MobileNo: mobile,
	}
}
func main() {
	nb := NewBike("Jatin", "Bangalore", "Classic 500", 9090909090)
	nb.Booked()
}
