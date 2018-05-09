package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	countdown(&buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
