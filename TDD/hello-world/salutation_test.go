package main

import "testing"

func TestSalutation(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' ", got, want)
		}
	}

	t.Run("salutation should be Hola in spanish", func(t *testing.T) {
		got := salutation("spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("salutation should be Bonjour in french", func(t *testing.T) {
		got := salutation("french")
		want := "Bonjour, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("salutation should be Hello if no language is specified", func(t *testing.T) {
		got := salutation("")
		want := "Hello, "
		assertCorrectMessage(t, got, want)
	})
}
