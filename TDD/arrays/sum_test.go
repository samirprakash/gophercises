package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {

	t.Run("return a slice containing sum of slices passed as params", func(t *testing.T) {
		got := sumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
