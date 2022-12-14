package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sezal", "English")
		want := "Hello, Sezal"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello In Spanish", func(t *testing.T) {
		got := Hello("Jim", "Spanish")
		want := "Hola, Jim"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Sezal", "French")
		want := "Bonjour, Sezal"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in Hindi", func(t *testing.T) {
		got := Hello("Sezal", "Hindi")
		want := "Namaste, Sezal"
		assertCorrectMessage(t, got, want)
	})
}

// refracted duplicate code in both tests into a func, to reuse
func assertCorrectMessage(t testing.TB, got, want string) {
	// it lets the 'test suite' know  this func is 'helper func'
	t.Helper()
	if got != want {
		t.Errorf("got : %q , want : %q", got, want)
	}
}
