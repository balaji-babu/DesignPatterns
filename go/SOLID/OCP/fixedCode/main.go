package main

import "fmt"

type (
	Color  int
	Size   int
	Rating int
)

const (
	red Color = iota
	blue
	green
	violet
	purple
)

const (
	small Size = iota
	medium
	large
	xl
	xxl
)

const (
	oneStar Rating = iota
	twoStar
	threeStar
	fourStar
	fiveStar
)

type Product struct {
	name   string
	color  Color
	size   Size
	rating Rating
}
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type Filter struct{}

func (f *Filter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	shirt := Product{"shirt", blue, xl, fourStar}
	tees := Product{"tees", green, large, fiveStar}
	jeans := Product{"jeans", blue, large, threeStar}

	products := []Product{shirt, tees, jeans}
	blueSpec := ColorSpecification{blue}
	bf := Filter{}
	for _, v := range bf.Filter(products, blueSpec) {
		fmt.Printf(" - %s is blue\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}
}
