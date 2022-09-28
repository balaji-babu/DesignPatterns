// To organize the DeepCoping we can have a function, which does the deep copying job.
package main

import "fmt"

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country}
}

func (p *Person) DeepCopy() *Person {
	q := *p // copies Name
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	p1 := Person{
		Name:    "Michael",
		Address: &Address{StreetAddress: "33, Thomaiyar Street, Kamarajapuram, pammal", City: "Chennai", Country: "India"},
		Friends: []string{"Alisa", "Mark"},
	}

	p2 := p1.DeepCopy()
	p2.Name = "Mathew"
	p2.Address.StreetAddress = "02, Thomaiyar Street, Kamarajapuram, pammal"
	p2.Friends = []string{"Mike", "Farooq"}

	fmt.Printf("P1 Details: \n\t Name: %s, \n\t Address:%#v \n\t Friends: %#v\n", p1.Name, p1.Address, p1.Friends)
	fmt.Printf("P2 Details: \n\t Name: %s, \n\t Address:%#v \n\t Friends: %#v\n", p2.Name, p2.Address, p2.Friends)

}
