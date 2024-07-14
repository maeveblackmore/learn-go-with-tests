package string

import (
	"fmt"
	"testing"
)

func TestStripWhitespace(t *testing.T) {
	string := StripWhitespace(" this should be stripped!  ")
	expected := "this should be stripped!"

	if string != expected {
		t.Errorf("expected '%s' but received '%s'", string, expected)
	}
}

func ExampleStripWhitespace() {
	string := StripWhitespace(" Hello, world! ")
	fmt.Println(string)
	// Output: Hello, world!
}
