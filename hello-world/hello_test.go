package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to a person", func(t *testing.T) {
		got := Hello("Sarah", "")
		want := "Hello Sarah !"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to everyone", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World !"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Deutch", func(t *testing.T) {
		got := Hello("Elodie", "Deutch")
		want := "Hallo Elodie !"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
