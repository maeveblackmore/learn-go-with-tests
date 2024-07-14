package string

import "strings"

// StripWhitespace returns the input with all leading and trailing whitespace removed.
func StripWhitespace(input string) string {
	return strings.TrimSpace(input)
}
