package main

import "testing"

func TestSum(t *testing.T) {

	assertCorrectAnswer := func(t *testing.T, n []int, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, n)
		}
	}

	t.Run("sum a collection of 5 numbers", func(t *testing.T) {
		n := []int{1, 2, 3, 4, 5}
		got := sum(n)
		want := 15
		assertCorrectAnswer(t, n, got, want)
	})

	t.Run("sum a collection of any size", func(t *testing.T) {
		n := []int{1, 2, 3}
		got := sum(n)
		want := 6
		assertCorrectAnswer(t, n, got, want)
	})
}
