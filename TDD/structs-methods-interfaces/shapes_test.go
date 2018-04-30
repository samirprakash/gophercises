package main

import "testing"

func TestShapes(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("calculate area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("calculate area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
