package main

import "testing"

func TestSlice(t *testing.T) {

	t.Run("return a slice of specified length", func(t *testing.T) {
		got := slice(5)
		want := []int{0, 0, 0, 0, 0}

		if got == nil {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
