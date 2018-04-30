package main

import "math"

// Rectangle defines the dimensions of a rectangle
type Rectangle struct {
	Length  float64
	Breadth float64
}

// Circle defines the dimensions of a circle
type Circle struct {
	Radius float64
}

func (r Rectangle) perimeter() (p float64) {
	p = 2 * (r.Length + r.Breadth)
	return
}

func (r Rectangle) area() (a float64) {
	a = r.Length * r.Breadth
	return
}

func (c Circle) area() (a float64) {
	a = math.Pi * c.Radius * c.Radius
	return
}
