package main

import "fmt"

type Document struct {
	data string
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
	color string
}

func (m MultiFunctionPrinter) Print(d Document) {
	fmt.Printf("Print: '%s' in %s format.\n", d.data, m.color)
}

func (m MultiFunctionPrinter) Fax(d Document) {
	fmt.Printf("Fax: '%s' in %s format.\n", d.data, m.color)
}

func (m MultiFunctionPrinter) Scan(d Document) {
	fmt.Printf("Scan: '%s' in %s format.\n", d.data, m.color)
}

type OldFashionedPrinter struct {
	color string
}

func (o OldFashionedPrinter) Print(d Document) {
	fmt.Printf("Print: '%s' in %s format.\n", d.data, o.color)
}

// Deprecated
func (o OldFashionedPrinter) Fax(d Document) {
	fmt.Println("Deprecated: Operation Not supported")
}

// Deprecated
func (o OldFashionedPrinter) Scan(d Document) {
	fmt.Println("Deprecated: Operation Not supported")
}

func PrinterFunction(m Machine, data Document) {
	fmt.Println("\n")
	m.Print(data)
	m.Fax(data)
	m.Scan(data)
}
func main() {
	m := MultiFunctionPrinter{"Color"}
	data1 := Document{"This is a MultiFunctionPrinter Document Data"}
	PrinterFunction(m, data1)

	// This Breaks ISP principle, Even though old format printer doesn't support fax and scan, we are forced to implement it.
	o := OldFashionedPrinter{"B&W"}
	data2 := Document{"This is a OldFunctionPrinter Document Data"}
	PrinterFunction(o, data2)
}
