package greeting

import "testing"

func TestGreet(t *testing.T) {
	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Greet("en", "Maeve")
		want := "Hello, Maeve!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in russian", func(t *testing.T) {
		got := Greet("ru", "Maeve")
		want := "привет, Maeve!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in japanese", func(t *testing.T) {
		got := Greet("jp", "Maeve")
		want := "こんにちは, Maeve!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Greet("en", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("providing an unsupported language", func(t *testing.T) {
		got := Greet("zz", "")
		want := "Unsupported language"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // Reports line number inside the function call, rather than this test helper.
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
