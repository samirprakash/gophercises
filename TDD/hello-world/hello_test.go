package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' ", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Sam")
		want := "Hello, Sam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello world when an empty string is supplied", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
