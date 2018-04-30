package main

import "math"

// Shape contains area() to be implemented
type Shape interface {
	area() float64
}

// Rectangle defines the dimensions of a rectangle
type Rectangle struct {
	Length  float64
	Breadth float64
}

// Triangle defines the dimensions of a triangle
type Triangle struct {
	Base   float64
	Height float64
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

func (t Triangle) area() (a float64) {
	a = (t.Base * t.Height) / 2
	return
}
