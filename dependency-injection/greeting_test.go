package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Create an empty buffer.
	buffer := bytes.Buffer{}

	// Pass a pointer to our buffer so it can be mutated.
	// Greet will modify this buffer so it's "Hello, <name>", where <name> is the second argument.
	Greet(&buffer, "Chris")

	// Convert our buffer to a string.
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
