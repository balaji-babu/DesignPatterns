package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var count = 0

type Book struct {
	titles []string
}

func (b *Book) String() string {
	return strings.Join(b.titles, "\n")
}

func (b *Book) addTitle(name string) {
	count++
	entry := fmt.Sprintf("%d: %s", count, name)
	b.titles = append(b.titles, entry)
}

func (b *Book) removeTitle(index int) {
	b.titles = append(b.titles[:index], b.titles[index+1:]...)
}

// This breaks the concept of Single Responsible Principle
// In the future I might need to save the data with different line breaks, this is tightly coupled this will break the concept of SRP
func (b *Book) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(b.String()), 0644)
}

func main() {
	var b Book
	b.addTitle("The Intelligent Investor")
	b.addTitle("Rich DAD POOR DAD")
	b.addTitle("COMMON STOCKS AND UNCOMMON PROFITS")
	b.addTitle("The Psychology of Money")
	b.addTitle("Beating the street")

	b.removeTitle(2)

	b.Save("/tmp/test.txt")
}
