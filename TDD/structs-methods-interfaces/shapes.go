package main

// Rectangle defines the dimensions
type Rectangle struct {
	Length  float64
	Breadth float64
}

func perimeter(rectangle Rectangle) (p float64) {
	p = 2 * (rectangle.Length + rectangle.Breadth)
	return
}

func area(rectangle Rectangle) (a float64) {
	a = rectangle.Length * rectangle.Breadth
	return
}
