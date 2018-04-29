package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Sam")
	want := "Hello, Sam"

	if got != want {
		t.Errorf("got '%s' want '%s' ", got, want)
	}
}
