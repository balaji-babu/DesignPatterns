// Deep Copy: Copying the entire object, unlike pointing to the references.
package main

import "fmt"

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	StreetAddress, City, Country string
}

func main() {
	p1 := Person{
		Name:    "Michael",
		Address: &Address{StreetAddress: "33, Thomaiyar Street, Kamarajapuram, pammal", City: "Chennai", Country: "India"},
	}

	// Copy Michael data to Samengles.

	p2 := p1
	// Replace Michael data with samengles.
	p2.Name = "Samengles"
	// Consider they belong to the same street and same address except the door no

	//Uncomment the below line to see the shallow copy example
	//p2.Address.StreetAddress = "39, Thomaiyar Street, Kamarajapuram, pammal"

	//Lets print the output.
	// Here the p2 address is actually a reference to p1 address, if we make any changes in p2 then p1 will be updated.

	//Uncomment the below 2 lines to see the shallow copy example
	//fmt.Printf("P1 Details: \n\t Name: %s, \n\t Address:%#v \n", p1.Name, p1.Address)
	//fmt.Printf("P2 Details: \n\t Name: %s, \n\t Address:%#v \n", p2.Name, p2.Address)

	// To fix the issue, now we are going to deep copy the Address field.
	p3 := p1
	p3.Name = "Mathew"
	p3.Address = &Address{
		p1.Address.StreetAddress,
		p1.Address.City,
		p1.Address.Country,
	}
	p3.Address.StreetAddress = "02, Thomaiyar Street, Kamarajapuram, pammal"
	fmt.Printf("P1 Details: \n\t Name: %s, \n\t Address:%#v \n", p1.Name, p1.Address)
	fmt.Printf("P3 Details: \n\t Name: %s, \n\t Address:%#v \n", p3.Name, p3.Address)

}
