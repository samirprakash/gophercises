package integers

import "testing"

func TestAddr(t *testing.T) {
	got := add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
