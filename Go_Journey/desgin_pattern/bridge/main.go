package main

import "fmt"

// U have Circle, Square
// U have Raster, vector to display shape => So u have a lot of functions like: RasterCirle, RasterSquare, VectorCircle, VectorSquare
type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRender struct {
	//
}

func (v *VectorRender) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRender struct {
	//
}

func (r *RasterRender) RenderCircle(radius float32) {
	fmt.Println("Drawing a pixel for the circle of radius", radius)
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
	// render different types through renderer
	raster := RasterRender{}
	//vector := VectorRender{}
	circle := NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
