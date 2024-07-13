package greeting

import "fmt"

const (
	englishHelloPrefix  = "Hello"
	russianHelloPrefix  = "привет"
	japaneseHelloPrefix = "こんにちは"
)

// Takes the ISO 639-1 code of a language as the first argument, and returns a greeting string using the string given as the second argument.
// If an empty string is provided, it uses the string "world" as a placeholder.
func Greet(prefix, name string) string {
	if name == "" {
		name = "world"
	}

	switch prefix {
	case "en":
		prefix = englishHelloPrefix
	case "ru":
		prefix = russianHelloPrefix
	case "jp":
		prefix = japaneseHelloPrefix
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}
