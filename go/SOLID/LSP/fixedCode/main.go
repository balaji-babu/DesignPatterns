package main

import "fmt"

type sized interface {
	getWidth() int
	setWidth(width int)
	getHeight() int
	setHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) getWidth() int {
	return r.width
}

func (r *Rectangle) setWidth(width int) {
	r.width = width
}

func (r *Rectangle) getHeight() int {
	return r.height
}

func (r *Rectangle) setHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) getWidth() int {
	return s.width
}

func (s *Square) setWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) getHeight() int {
	return s.width
}

func (s *Square) setHeight(height int) {
	s.width = height
	s.height = height
}

// Fixed code
type Square2 struct {
	size int
}

func (s *Square2) Rectangle() *Rectangle {
	return &Rectangle{s.size, s.size}
}

func calculate(sized sized) {
	width := sized.getWidth()
	sized.setHeight(10)
	expectedArea := 10 * width
	actualArea := sized.getWidth() * sized.getHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{4, 5}
	calculate(rc)

	sq1 := NewSquare(5)
	calculate(sq1)

	sq2 := Square2{5}
	calculate(sq2.Rectangle())

}
