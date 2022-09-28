// To organize the DeepCoping we can have a function, which does the deep copying job.
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Employee struct {
	Name   string
	Office Address
}

type Address struct {
	StreetAddress, City, Country string
	Suite                        int
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{"", Address{"Kota Jakarta Selatan, Daerah Khusus Ibukota", "Jakarta", "Indonesia", 0}}
var auxOffice = Employee{"", Address{"Diamond District, 2nd, 4th & 6th Floor Tower B, HAL Old Airport Rd,", "Bengaluru", "India", 0}}

func newEmployee(proto *Employee,
	name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	p1 := NewMainOfficeEmployee("Michael", 100)
	p2 := NewAuxOfficeEmployee("Mathew", 200)

	fmt.Printf("P1 Details: %#v \n", p1)
	fmt.Printf("P2 Details: %#v \n", p2)

}
