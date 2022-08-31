package main

import "fmt"

type Document struct {
	data string
}

// Split in to several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Fax interface {
	Fax(d Document)
}

type MultiFunctionPrinterInterface interface {
	Printer
	Scanner
	Fax
}

type OldFunctionPrinterInterface interface {
	Printer
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

func main() {
	var m MultiFunctionPrinterInterface
	m = MultiFunctionPrinter{"Color"}
	data1 := Document{"This is a MultiFunctionPrinter Document Data"}
	m.Print(data1)
	m.Fax(data1)
	m.Scan(data1)

	fmt.Println("========================================================================")
	var o OldFunctionPrinterInterface
	o = OldFashionedPrinter{"B&W"}
	data2 := Document{"This is a OldFunctionPrinter Document Data"}
	o.Print(data2)
}
