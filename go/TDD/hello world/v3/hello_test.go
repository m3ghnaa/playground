package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Meghna")
	want := "Hello, Meghna"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
