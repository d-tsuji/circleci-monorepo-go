package service1

import "testing"

func TestAdd(t *testing.T) {
	a, b := 1, 2
	want := 3
	if got := Add(a, b); got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}
