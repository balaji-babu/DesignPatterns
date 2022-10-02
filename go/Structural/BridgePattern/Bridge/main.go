package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

type RasterRenderer struct {
	Dpi int
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (v *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}
func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	Circle1 := NewCircle(&raster, 5)
	Circle1.Draw()
	Circle1.Resize(2)
	Circle1.Draw()

	Circle2 := NewCircle(&vector, 5)
	Circle2.Draw()
	Circle2.Resize(4)
	Circle2.Draw()
}
