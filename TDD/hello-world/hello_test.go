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
		got := hello("Sam", "")
		want := "Hello, Sam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello world when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := hello("Sam", "spanish")
		want := "Hola, Sam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := hello("Sam", "french")
		want := "Bonjour, Sam"
		assertCorrectMessage(t, got, want)
	})
}

func TestSalutation(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' ", got, want)
		}
	}

	t.Run("Salutation should be Hola in spanish", func(t *testing.T) {
		got := salutation("spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("Salutation should be Bonjour in french", func(t *testing.T) {
		got := salutation("french")
		want := "Bonjour, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("Salutation should be Hello if no language is specified", func(t *testing.T) {
		got := salutation("")
		want := "Hello, "
		assertCorrectMessage(t, got, want)
	})
}
