package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency * 100.0)
}

func main() {

	circle := Circle{2}
	circle.Resize(7)
	fmt.Println(circle.Render())

	square := Square{5}
	fmt.Println(square.Render())

	blueCircle := ColoredShape{&circle, "blue"}
	fmt.Println(blueCircle.Render())

	redSquare := ColoredShape{&square, "red"}
	fmt.Println(redSquare.Render())

	// redSquare.Resize(7) not available on the ColoredShape

	// we can now apply the ColoredShape decorator over another decorator
	bhsCircle := TransparentShape{&blueCircle, 0.5}
	fmt.Println(bhsCircle.Render())

}
