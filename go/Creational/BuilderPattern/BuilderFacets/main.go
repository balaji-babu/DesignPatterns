package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Address, Pincode, City string
	CompanyName, Position string
	AnnualIncome int 
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder(firstName, lastName string) *PersonBuilder {
	return &PersonBuilder{
		&Person {
			FirstName: firstName,
			LastName: lastName,
		},
	}
} 

func (pb *PersonBuilder) Build() *Person {
    return pb.person
}

type PersonJobBuilder struct{
	PersonBuilder
}


func (pb *PersonBuilder) Works() *PersonJobBuilder{
	return &PersonJobBuilder{*pb}
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
  pjb.person.CompanyName = companyName
  return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
  pjb.person.Position = position
  return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
  pjb.person.AnnualIncome = annualIncome
  return pjb
}

type PersonAddressBuilder struct{
	PersonBuilder
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
  return &PersonAddressBuilder{*pb}
}


func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
  pab.person.Address = streetAddress
  return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
  pab.person.City = city
  return pab
}

func (pab *PersonAddressBuilder) WithPincode(pincode string) *PersonAddressBuilder {
  pab.person.Pincode = pincode
  return pab
}


func main() {
	pb := NewPersonBuilder("Rahul","singh")
	person := pb.Build()

	fmt.Println("Person: ", *person)

	person = pb.Lives().At("Brigade Road").In("Bangalore").WithPincode("560001").Works().At("ABC").AsA("Senior DataScience Engineer").Earning(5000000).Build()
	fmt.Printf("PersonInfo: %#v", *person)
}