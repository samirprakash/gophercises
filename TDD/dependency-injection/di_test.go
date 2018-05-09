package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	greet(&buffer, "Sam")

	got := buffer.String()
	want := "Hello, Sam"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
