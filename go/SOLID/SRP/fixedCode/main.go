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

func (b *Book) addTitle(name string) {
	count++
	entry := fmt.Sprintf("%d: %s", count, name)
	b.titles = append(b.titles, entry)
}

func (b *Book) removeTitle(index int) {
	b.titles = append(b.titles[:index], b.titles[index+1:]...)
}

type APersistence struct {
	lineSeparator string
}

func (ap *APersistence) Save(b *Book, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(b.titles, ap.lineSeparator)), 0644)
}

func main() {
	var b Book
	b.addTitle("The Intelligent Investor")
	b.addTitle("Rich DAD POOR DAD")
	b.addTitle("COMMON STOCKS AND UNCOMMON PROFITS")
	b.addTitle("The Psychology of Money")
	b.addTitle("Beating the street")

	b.removeTitle(2)

	ap := APersistence{"\n"}
	ap.Save(&b, "/tmp")
}
