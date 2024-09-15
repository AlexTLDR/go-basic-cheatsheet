package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}
func main() {
	rect := Rectangle{Width: 5, Height: 4}
	circle := Circle{Radius: 2}

	fmt.Println("Rectangle area:", calculateArea(rect))
	fmt.Println("Circle area:", calculateArea(circle))
}
