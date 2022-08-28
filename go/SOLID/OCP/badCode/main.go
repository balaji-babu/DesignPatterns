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

type Filter struct {
}

func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// In we need to filter based on size and rating, we need to modify the code, which leads to modification of code
func main() {
	shirt := Product{"shirt", blue, xl, fourStar}
	tees := Product{"tees", green, large, fiveStar}
	jeans := Product{"jeans", blue, large, threeStar}

	products := []Product{shirt, tees, jeans}
	f := Filter{}
	for _, v := range f.filterByColor(products, blue) {
		fmt.Printf(" - %s is blue\n", v.name)
	}
}
