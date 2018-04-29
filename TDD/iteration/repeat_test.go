package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	want := "aaaaaaa"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	o := Repeat("a", 7)
	fmt.Println(o)
	// Output: aaaaaaa
}
