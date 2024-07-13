package greeting

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Maeve")
	want := "Hello, Maeve!"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}
