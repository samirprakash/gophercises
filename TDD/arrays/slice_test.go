package main

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {

	t.Run("return a slice of specified length", func(t *testing.T) {
		got := slice(5)
		want := []int{0, 0, 0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
