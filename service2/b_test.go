package service2

import "testing"

func TestSub(t *testing.T) {
	a, b := 1, 2
	want := -1
	if got := Sub(a, b); got != want {
		t.Errorf("Sub() = %v, want %v", got, want)
	}
}
