package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return r.width*2 + r.height*2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
