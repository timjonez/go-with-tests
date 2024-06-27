package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Tim", "")
		want := "Hello, Tim"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say Hello default to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in spanish", func(t *testing.T) {
		got := Hello("John", "Spanish")
		want := "Hola, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in french", func(t *testing.T) {
		got := Hello("Tim", "French")
		want := "Bonjour, Tim"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in russian", func(t *testing.T) {
		got := Hello("Tim", "Russian")
		want := "Привет, Tim"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
