package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Sam")
		want := "Hello, Sam"

		if got != want {
			t.Errorf("got '%s' want '%s' ", got, want)
		}
	})

	t.Run("Saying hello world when an empty string is supplied", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

}
