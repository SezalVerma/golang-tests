package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sezal")
	want := "Hello, Sezal"

	if got != want {
		t.Errorf("got : %q , want : %q", got, want)
	}
}
