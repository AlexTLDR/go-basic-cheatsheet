package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())

}

func main() {
	rect := Rectangle{Width: 5, Height: 4}
	describeShape(rect)
	val := 25.0
	result, err := performCalculation(val)
	fmt.Printf("Square root of %v is %v and error is %v\n", val, result, err)
}

type CalculationError struct {
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{"Value cannot be negative"}
	}
	return math.Sqrt(val), nil
}
